package main

import (
	"fmt"
	"log"
	"time"

	"github.com/James2898/alumni-platform/backend/internal/config"
	"github.com/James2898/alumni-platform/backend/internal/domain"
	"github.com/James2898/alumni-platform/backend/internal/handler"
	"github.com/James2898/alumni-platform/backend/internal/middleware"
	"github.com/James2898/alumni-platform/backend/internal/repository"
	"github.com/James2898/alumni-platform/backend/internal/service"
	"github.com/James2898/alumni-platform/backend/internal/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	cfg := config.LoadConfig()

	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 1. Connect to MySQL Database with retry loop for Docker cold start
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var db *gorm.DB
	var err error

	for attempts := 1; attempts <= 15; attempts++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			sqlDB, pingErr := db.DB()
			if pingErr == nil && sqlDB.Ping() == nil {
				log.Println("Successfully connected to MySQL database!")
				break
			}
		}
		log.Printf("Waiting for database connection (attempt %d/15)...", attempts)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Fatal: could not establish database connection: %v", err)
	}

	// 2. Initialize MinIO / S3 Storage Service
	storageSvc, err := storage.NewMinioStorageService(cfg)
	if err != nil {
		log.Printf("Warning: MinIO storage client failed to initialize: %v", err)
	} else {
		log.Println("MinIO storage service initialized successfully.")
	}

	// 3. Initialize Repositories
	userRepo := repository.NewUserRepository(db)
	alumniRepo := repository.NewAlumniRepository(db)
	courseRepo := repository.NewCourseRepository(db)

	// 4. Initialize Services
	authSvc := service.NewAuthService(userRepo, alumniRepo, courseRepo, cfg)
	alumniSvc := service.NewAlumniService(alumniRepo, storageSvc)
	courseSvc := service.NewCourseService(courseRepo)

	// 5. Initialize Handlers
	healthHandler := handler.NewHealthHandler(db)
	authHandler := handler.NewAuthHandler(authSvc)
	alumniHandler := handler.NewAlumniHandler(alumniSvc)
	courseHandler := handler.NewCourseHandler(courseSvc)
	adminHandler := handler.NewAdminHandler(userRepo, courseSvc)

	// 6. Router Setup
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// Native Cloud / Kubernetes Health Probes (CKAD pattern)
	r.GET("/healthz", healthHandler.Healthz)
	r.GET("/readyz", healthHandler.Readyz)

	api := r.Group("/api")
	{
		// Public Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/logout", authHandler.Logout)
			auth.GET("/me", middleware.AuthMiddleware(authSvc), authHandler.Me)
		}

		// Public Course catalog
		api.GET("/courses", courseHandler.GetAll)

		// Public Alumni Directory & Details
		api.GET("/alumni", alumniHandler.GetDirectory)
		api.GET("/alumni/:id", alumniHandler.GetByID)

		// Protected Alumni Self-Service Routes
		alumniProtected := api.Group("/alumni")
		alumniProtected.Use(middleware.AuthMiddleware(authSvc))
		{
			alumniProtected.GET("/profile/me", alumniHandler.GetMyProfile)
			alumniProtected.PUT("/profile/me", alumniHandler.UpdateMyProfile)
			alumniProtected.POST("/upload-ticket", alumniHandler.GenerateUploadURL)
		}

		// Protected Admin Dashboard Routes
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(authSvc), middleware.RequireRole(domain.RoleAdmin))
		{
			admin.GET("/pending-users", adminHandler.ListPendingUsers)
			admin.PATCH("/users/:id/status", adminHandler.UpdateUserStatus)
			admin.POST("/courses", adminHandler.CreateCourse)
			admin.DELETE("/courses/:id", adminHandler.DeleteCourse)
		}
	}

	serverAddr := ":" + cfg.Port
	log.Printf("Alumni Platform REST API server listening on %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
}

package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/James2898/alumni-platform/backend/internal/config"
	"github.com/James2898/alumni-platform/backend/internal/domain"
	"github.com/James2898/alumni-platform/backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CustomClaims struct {
	UserID string          `json:"user_id"`
	Email  string          `json:"email"`
	Role   domain.UserRole `json:"role"`
	jwt.RegisteredClaims
}

type AuthService interface {
	RegisterAlumni(ctx context.Context, req domain.RegisterRequest) (*domain.User, error)
	Login(ctx context.Context, email, password string) (*domain.AuthResponse, error)
	GenerateTokenPair(user *domain.User) (accessToken, refreshToken string, err error)
	ValidateToken(tokenStr string) (*CustomClaims, error)
	GetCurrentUser(ctx context.Context, userID string) (*domain.User, error)
}

type authService struct {
	userRepo   repository.UserRepository
	alumniRepo repository.AlumniRepository
	courseRepo repository.CourseRepository
	cfg        *config.Config
}

func NewAuthService(
	userRepo repository.UserRepository,
	alumniRepo repository.AlumniRepository,
	courseRepo repository.CourseRepository,
	cfg *config.Config,
) AuthService {
	return &authService{
		userRepo:   userRepo,
		alumniRepo: alumniRepo,
		courseRepo: courseRepo,
		cfg:        cfg,
	}
}

func (s *authService) RegisterAlumni(ctx context.Context, req domain.RegisterRequest) (*domain.User, error) {
	// Verify course exists
	_, err := s.courseRepo.FindByID(ctx, req.CourseID)
	if err != nil {
		return nil, errors.New("selected course does not exist")
	}

	// Check if email already registered
	existing, _ := s.userRepo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New("email is already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	userID := uuid.New().String()
	user := &domain.User{
		ID:           userID,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         domain.RoleAlumni,
		Status:       domain.StatusPending, // New registrations require admin approval
	}

	profile := &domain.AlumniProfile{
		ID:             uuid.New().String(),
		UserID:         userID,
		StudentNumber:  req.StudentNumber,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		CourseID:       req.CourseID,
		GraduationYear: req.GraduationYear,
		Phone:          req.Phone,
	}

	// Execute inside database transaction
	err = s.userRepo.DB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.userRepo.CreateUser(ctx, tx, user); err != nil {
			return err
		}
		if err := s.alumniRepo.CreateProfile(ctx, tx, profile); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("registration failed: %w", err)
	}

	return s.userRepo.FindByID(ctx, userID)
}

func (s *authService) Login(ctx context.Context, email, password string) (*domain.AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	if user.Status == domain.StatusPending {
		return nil, errors.New("your account is pending administrator approval")
	}
	if user.Status == domain.StatusSuspended {
		return nil, errors.New("your account has been suspended")
	}

	accessToken, refreshToken, err := s.GenerateTokenPair(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	return &domain.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *user,
	}, nil
}

func (s *authService) GenerateTokenPair(user *domain.User) (string, string, error) {
	now := time.Now()

	accessClaims := CustomClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   user.ID,
		},
	}

	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := accessTokenObj.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return "", "", err
	}

	refreshClaims := CustomClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   user.ID,
		},
	}

	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refreshTokenObj.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *authService) ValidateToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(s.cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (s *authService) GetCurrentUser(ctx context.Context, userID string) (*domain.User, error) {
	return s.userRepo.FindByID(ctx, userID)
}

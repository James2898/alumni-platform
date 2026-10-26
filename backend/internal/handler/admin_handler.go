package handler

import (
	"net/http"

	"github.com/James2898/alumni-platform/backend/internal/domain"
	"github.com/James2898/alumni-platform/backend/internal/repository"
	"github.com/James2898/alumni-platform/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	userRepo  repository.UserRepository
	courseSvc service.CourseService
}

func NewAdminHandler(userRepo repository.UserRepository, courseSvc service.CourseService) *AdminHandler {
	return &AdminHandler{
		userRepo:  userRepo,
		courseSvc: courseSvc,
	}
}

func (h *AdminHandler) ListPendingUsers(c *gin.Context) {
	users, err := h.userRepo.ListPendingAlumniUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list pending users: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *AdminHandler) UpdateUserStatus(c *gin.Context) {
	id := c.Param("id")

	var req domain.UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userRepo.UpdateStatus(c.Request.Context(), id, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user status: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user status updated successfully", "status": req.Status})
}

func (h *AdminHandler) CreateCourse(c *gin.Context) {
	var req domain.CourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.courseSvc.CreateCourse(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create course: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, course)
}

func (h *AdminHandler) DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	if err := h.courseSvc.DeleteCourse(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete course: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "course deleted successfully"})
}

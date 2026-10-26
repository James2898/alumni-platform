package handler

import (
	"net/http"

	"github.com/James2898/alumni-platform/backend/internal/domain"
	"github.com/James2898/alumni-platform/backend/internal/middleware"
	"github.com/James2898/alumni-platform/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AlumniHandler struct {
	alumniSvc service.AlumniService
}

func NewAlumniHandler(alumniSvc service.AlumniService) *AlumniHandler {
	return &AlumniHandler{alumniSvc: alumniSvc}
}

func (h *AlumniHandler) GetDirectory(c *gin.Context) {
	var filter domain.AlumniFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.alumniSvc.GetDirectory(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch directory: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *AlumniHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	profile, err := h.alumniSvc.GetProfileByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "alumni profile not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *AlumniHandler) GetMyProfile(c *gin.Context) {
	userIDVal, _ := c.Get(middleware.ContextUserIDKey)
	userID, _ := userIDVal.(string)

	profile, err := h.alumniSvc.GetProfileByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "profile not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *AlumniHandler) UpdateMyProfile(c *gin.Context) {
	userIDVal, _ := c.Get(middleware.ContextUserIDKey)
	userID, _ := userIDVal.(string)

	var req domain.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.alumniSvc.UpdateProfile(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update profile: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *AlumniHandler) GenerateUploadURL(c *gin.Context) {
	userIDVal, _ := c.Get(middleware.ContextUserIDKey)
	userID, _ := userIDVal.(string)

	var req domain.PresignedURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.alumniSvc.CreatePresignedUpload(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate upload ticket: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

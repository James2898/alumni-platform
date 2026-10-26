package handler

import (
	"net/http"

	"github.com/James2898/alumni-platform/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseSvc service.CourseService
}

func NewCourseHandler(courseSvc service.CourseService) *CourseHandler {
	return &CourseHandler{courseSvc: courseSvc}
}

func (h *CourseHandler) GetAll(c *gin.Context) {
	courses, err := h.courseSvc.GetAllCourses(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve courses: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

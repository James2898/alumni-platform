package service

import (
	"context"

	"github.com/James2898/alumni-platform/backend/internal/domain"
	"github.com/James2898/alumni-platform/backend/internal/repository"
	"github.com/google/uuid"
)

type CourseService interface {
	GetAllCourses(ctx context.Context) ([]domain.Course, error)
	CreateCourse(ctx context.Context, req domain.CourseRequest) (*domain.Course, error)
	DeleteCourse(ctx context.Context, id string) error
}

type courseService struct {
	courseRepo repository.CourseRepository
}

func NewCourseService(courseRepo repository.CourseRepository) CourseService {
	return &courseService{courseRepo: courseRepo}
}

func (s *courseService) GetAllCourses(ctx context.Context) ([]domain.Course, error) {
	return s.courseRepo.FindAll(ctx)
}

func (s *courseService) CreateCourse(ctx context.Context, req domain.CourseRequest) (*domain.Course, error) {
	course := &domain.Course{
		ID:         uuid.New().String(),
		Code:       req.Code,
		Name:       req.Name,
		Department: req.Department,
	}
	if err := s.courseRepo.Create(ctx, course); err != nil {
		return nil, err
	}
	return course, nil
}

func (s *courseService) DeleteCourse(ctx context.Context, id string) error {
	return s.courseRepo.Delete(ctx, id)
}

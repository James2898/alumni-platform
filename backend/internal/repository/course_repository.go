package repository

import (
	"context"

	"github.com/James2898/alumni-platform/backend/internal/domain"
	"gorm.io/gorm"
)

type CourseRepository interface {
	FindAll(ctx context.Context) ([]domain.Course, error)
	FindByID(ctx context.Context, id string) (*domain.Course, error)
	Create(ctx context.Context, course *domain.Course) error
	Delete(ctx context.Context, id string) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) FindAll(ctx context.Context) ([]domain.Course, error) {
	var courses []domain.Course
	err := r.db.WithContext(ctx).Order("code asc").Find(&courses).Error
	return courses, err
}

func (r *courseRepository) FindByID(ctx context.Context, id string) (*domain.Course, error) {
	var course domain.Course
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&course).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *courseRepository) Create(ctx context.Context, course *domain.Course) error {
	return r.db.WithContext(ctx).Create(course).Error
}

func (r *courseRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Course{}).Error
}

package repository

import (
	"context"

	"github.com/James2898/alumni-platform/backend/internal/domain"
	"gorm.io/gorm"
)

type AlumniRepository interface {
	CreateProfile(ctx context.Context, tx *gorm.DB, profile *domain.AlumniProfile) error
	FindByUserID(ctx context.Context, userID string) (*domain.AlumniProfile, error)
	FindByID(ctx context.Context, id string) (*domain.AlumniProfile, error)
	Update(ctx context.Context, profile *domain.AlumniProfile) error
	ListFiltered(ctx context.Context, filter domain.AlumniFilter) ([]domain.AlumniProfile, int64, error)
}

type alumniRepository struct {
	db *gorm.DB
}

func NewAlumniRepository(db *gorm.DB) AlumniRepository {
	return &alumniRepository{db: db}
}

func (r *alumniRepository) CreateProfile(ctx context.Context, tx *gorm.DB, profile *domain.AlumniProfile) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	return db.WithContext(ctx).Create(profile).Error
}

func (r *alumniRepository) FindByUserID(ctx context.Context, userID string) (*domain.AlumniProfile, error) {
	var profile domain.AlumniProfile
	err := r.db.WithContext(ctx).
		Preload("Course").
		Preload("WorkExperiences").
		Preload("User").
		Where("user_id = ?", userID).
		First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *alumniRepository) FindByID(ctx context.Context, id string) (*domain.AlumniProfile, error) {
	var profile domain.AlumniProfile
	err := r.db.WithContext(ctx).
		Preload("Course").
		Preload("WorkExperiences").
		Preload("User").
		Where("id = ?", id).
		First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *alumniRepository) Update(ctx context.Context, profile *domain.AlumniProfile) error {
	return r.db.WithContext(ctx).Save(profile).Error
}

func (r *alumniRepository) ListFiltered(ctx context.Context, filter domain.AlumniFilter) ([]domain.AlumniProfile, int64, error) {
	var profiles []domain.AlumniProfile
	var total int64

	query := r.db.WithContext(ctx).
		Model(&domain.AlumniProfile{}).
		Joins("JOIN users ON users.id = alumni_profiles.user_id").
		Where("users.status = ?", domain.StatusActive).
		Preload("Course").
		Preload("WorkExperiences")

	if filter.CourseID != "" {
		query = query.Where("alumni_profiles.course_id = ?", filter.CourseID)
	}

	if filter.GraduationYear > 0 {
		query = query.Where("alumni_profiles.graduation_year = ?", filter.GraduationYear)
	}

	if filter.Search != "" {
		searchPattern := "%" + filter.Search + "%"
		query = query.Where("(alumni_profiles.first_name LIKE ? OR alumni_profiles.last_name LIKE ? OR alumni_profiles.student_number LIKE ?)",
			searchPattern, searchPattern, searchPattern)
	}

	if filter.Company != "" {
		query = query.Joins("LEFT JOIN work_experiences ON work_experiences.alumni_id = alumni_profiles.id").
			Where("work_experiences.company_name LIKE ?", "%"+filter.Company+"%")
	}

	if err := query.Distinct("alumni_profiles.id").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.PageSize <= 0 || filter.PageSize > 100 {
		filter.PageSize = 12
	}

	offset := (filter.Page - 1) * filter.PageSize
	err := query.Distinct().
		Order("alumni_profiles.created_at desc").
		Limit(filter.PageSize).
		Offset(offset).
		Find(&profiles).Error

	return profiles, total, err
}

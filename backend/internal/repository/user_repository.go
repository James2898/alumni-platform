package repository

import (
	"context"

	"github.com/James2898/alumni-platform/backend/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx *gorm.DB, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	UpdateStatus(ctx context.Context, id string, status domain.UserStatus) error
	ListPendingAlumniUsers(ctx context.Context) ([]domain.User, error)
	DB() *gorm.DB
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) DB() *gorm.DB {
	return r.db
}

func (r *userRepository) CreateUser(ctx context.Context, tx *gorm.DB, user *domain.User) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	return db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).
		Preload("AlumniProfile.Course").
		Preload("AdminProfile").
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).
		Preload("AlumniProfile.Course").
		Preload("AlumniProfile.WorkExperiences").
		Preload("AdminProfile").
		Where("id = ?", id).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateStatus(ctx context.Context, id string, status domain.UserStatus) error {
	return r.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", id).Update("status", status).Error
}

func (r *userRepository) ListPendingAlumniUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	err := r.db.WithContext(ctx).
		Preload("AlumniProfile.Course").
		Where("role = ? AND status = ?", domain.RoleAlumni, domain.StatusPending).
		Order("created_at desc").
		Find(&users).Error
	return users, err
}

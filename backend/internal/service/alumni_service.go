package service

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/James2898/alumni-platform/backend/internal/domain"
	"github.com/James2898/alumni-platform/backend/internal/repository"
	"github.com/James2898/alumni-platform/backend/internal/storage"
	"github.com/google/uuid"
)

type AlumniService interface {
	GetDirectory(ctx context.Context, filter domain.AlumniFilter) (*domain.PaginatedAlumniResponse, error)
	GetProfileByID(ctx context.Context, id string) (*domain.AlumniProfile, error)
	GetProfileByUserID(ctx context.Context, userID string) (*domain.AlumniProfile, error)
	UpdateProfile(ctx context.Context, userID string, req domain.UpdateProfileRequest) (*domain.AlumniProfile, error)
	CreatePresignedUpload(ctx context.Context, userID string, req domain.PresignedURLRequest) (*domain.PresignedURLResponse, error)
}

type alumniService struct {
	alumniRepo repository.AlumniRepository
	storageSvc storage.StorageService
}

func NewAlumniService(alumniRepo repository.AlumniRepository, storageSvc storage.StorageService) AlumniService {
	return &alumniService{
		alumniRepo: alumniRepo,
		storageSvc: storageSvc,
	}
}

func (s *alumniService) GetDirectory(ctx context.Context, filter domain.AlumniFilter) (*domain.PaginatedAlumniResponse, error) {
	profiles, total, err := s.alumniRepo.ListFiltered(ctx, filter)
	if err != nil {
		return nil, err
	}

	for i := range profiles {
		if profiles[i].AvatarKey != "" {
			profiles[i].AvatarURL = s.storageSvc.GetPublicURL(profiles[i].AvatarKey)
		}
	}

	pageSize := filter.PageSize
	if pageSize <= 0 {
		pageSize = 12
	}
	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))

	return &domain.PaginatedAlumniResponse{
		Data:       profiles,
		Total:      total,
		Page:       filter.Page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

func (s *alumniService) GetProfileByID(ctx context.Context, id string) (*domain.AlumniProfile, error) {
	profile, err := s.alumniRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if profile.AvatarKey != "" {
		profile.AvatarURL = s.storageSvc.GetPublicURL(profile.AvatarKey)
	}
	return profile, nil
}

func (s *alumniService) GetProfileByUserID(ctx context.Context, userID string) (*domain.AlumniProfile, error) {
	profile, err := s.alumniRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if profile.AvatarKey != "" {
		profile.AvatarURL = s.storageSvc.GetPublicURL(profile.AvatarKey)
	}
	return profile, nil
}

func (s *alumniService) UpdateProfile(ctx context.Context, userID string, req domain.UpdateProfileRequest) (*domain.AlumniProfile, error) {
	profile, err := s.alumniRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if req.FirstName != "" {
		profile.FirstName = req.FirstName
	}
	if req.LastName != "" {
		profile.LastName = req.LastName
	}
	if req.Phone != "" {
		profile.Phone = req.Phone
	}
	if req.LinkedinURL != "" {
		profile.LinkedinURL = req.LinkedinURL
	}
	if req.Bio != "" {
		profile.Bio = req.Bio
	}
	if req.AvatarKey != "" {
		profile.AvatarKey = req.AvatarKey
	}

	if err := s.alumniRepo.Update(ctx, profile); err != nil {
		return nil, err
	}

	return s.GetProfileByUserID(ctx, userID)
}

func (s *alumniService) CreatePresignedUpload(ctx context.Context, userID string, req domain.PresignedURLRequest) (*domain.PresignedURLResponse, error) {
	ext := filepath.Ext(req.FileName)
	randomID := uuid.New().String()
	objectKey := fmt.Sprintf("%s/%s/%s%s", req.Category, userID, randomID, ext)

	// 15-minute presigned upload window
	uploadURL, err := s.storageSvc.GeneratePresignedUploadURL(ctx, objectKey, req.ContentType, 15*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("failed to generate presigned url: %w", err)
	}

	return &domain.PresignedURLResponse{
		UploadURL: uploadURL,
		ObjectKey: objectKey,
		FileURI:   s.storageSvc.GetPublicURL(objectKey),
	}, nil
}

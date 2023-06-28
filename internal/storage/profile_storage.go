package storage

import (
	"context"
	"fmt"
	"github.com/MarySmirnova/tikkichest-profile-service/internal/model"
	"log"
)

type DBStorage interface {
	GetProfile(ctx context.Context, username string) (*model.Profile, error)
	PageOfProfiles(ctx context.Context, itemsPerPage uint, from string) ([]*model.Profile, string, error)
	CreateProfile(ctx context.Context, profile *model.Profile) error
	UpdateProfile(ctx context.Context, username string, profile *model.ProfileUpdate) error
	DeleteProfile(ctx context.Context, username string) error
}

type CacheStorage interface {
	GetProfile(ctx context.Context, username string) (*model.Profile, error)
	AddProfile(ctx context.Context, profile *model.Profile) error
	DeleteProfile(ctx context.Context, username string) (*model.Profile, error)
}

type ProfileStore struct {
	db    DBStorage
	cache CacheStorage
}

func NewStore(db DBStorage, cache CacheStorage) *ProfileStore {
	return &ProfileStore{
		db:    db,
		cache: cache,
	}
}

func (s *ProfileStore) GetProfile(ctx context.Context, username string) (*model.Profile, error) {
	profile, err := s.cache.GetProfile(ctx, username)
	if err == nil {
		return profile, nil
	}

	return s.db.GetProfile(ctx, username)
}

func (s *ProfileStore) PageOfProfiles(ctx context.Context, itemsPerPage uint, from string) ([]*model.Profile, string, error) {
	return s.db.PageOfProfiles(ctx, itemsPerPage, from)
}

func (s *ProfileStore) CreateProfile(ctx context.Context, profile *model.Profile) error {
	if err := s.db.CreateProfile(ctx, profile); err != nil {
		return err
	}

	if err := s.cache.AddProfile(ctx, profile); err != nil {
		log.Printf("failed to add profile in cache: %v", err)
	}

	return nil
}

func (s *ProfileStore) UpdateProfile(ctx context.Context, username string, profile *model.ProfileUpdate) error {
	oldProfile, err := s.cache.DeleteProfile(ctx, username)
	if err != nil {
		return fmt.Errorf("error with update profile, faled to invalide profile cache: %w", err)
	}

	if err := s.db.UpdateProfile(ctx, username, profile); err != nil {
		return err
	}

	oldProfile.Update(profile)

	if err := s.cache.AddProfile(ctx, oldProfile); err != nil {
		log.Printf("failed to add profile in cache: %v", err)
	}

	return nil
}

func (s *ProfileStore) DeleteProfile(ctx context.Context, username string) error {
	_, err := s.cache.DeleteProfile(ctx, username)
	if err != nil {
		return fmt.Errorf("error with delete profile, faled to invalide profile cache: %w", err)
	}

	return s.db.DeleteProfile(ctx, username)
}

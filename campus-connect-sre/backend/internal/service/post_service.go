package service

import (
	"context"

	"campus-connect/backend/internal/models"
)

type repository interface {
	ListPosts(ctx context.Context) ([]models.Post, error)
	CreatePost(ctx context.Context, post models.Post) (models.Post, error)
	LikePost(ctx context.Context, id int64) error
}

type PostService struct {
	repo repository
}

func NewPostService(repo repository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) ListPosts(ctx context.Context) ([]models.Post, error) {
	return s.repo.ListPosts(ctx)
}

func (s *PostService) CreatePost(ctx context.Context, post models.Post) (models.Post, error) {
	return s.repo.CreatePost(ctx, post)
}

func (s *PostService) LikePost(ctx context.Context, id int64) error {
	return s.repo.LikePost(ctx, id)
}

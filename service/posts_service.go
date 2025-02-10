package service

import (
	"article/models"
	"article/repository"
	"errors"
	"strings"
)

type PostService struct {
	postRepo *repository.PostsRepository
}

func NewPostsService(repo *repository.PostsRepository) *PostService {
	return &PostService{postRepo: repo}
}

func (s *PostService) Create(posts *models.Posts) error {
	if strings.TrimSpace(posts.Title) == "" {
		return errors.New("title is required")
	}

	if len(posts.Title) < 20 {
		return errors.New("title must be at least 20 characters")
	}

	if strings.TrimSpace(posts.Content) == "" {
		return errors.New("content is required")
	}

	if len(posts.Content) < 200 {
		return errors.New("content must be at least 200 characters")
	}

	if strings.TrimSpace(posts.Category) == "" {
		return errors.New("category is required")
	}

	if len(posts.Category) < 3 {
		return errors.New("category must be at least 3 characters")
	}

	if strings.TrimSpace(posts.Status) == "" {
		return errors.New("status is required")
	}

	if posts.Status != "publish" && posts.Status != "draft" && posts.Status != "thrash" {
		return errors.New("status must be 'publish', 'draft', or 'thrash'")
	}

	return s.postRepo.Create(posts)
}

func (s *PostService) GetPosts(limit int, offset int) ([]models.Posts, error) {
	return s.postRepo.GetPosts(limit, offset)
}

func (s *PostService) GetPostsById(id int) (*models.Posts, error) {
	return s.postRepo.GetPostsById(id)
}

func (s *PostService) Update(posts *models.Posts) error {
	if strings.TrimSpace(posts.Title) == "" {
		return errors.New("title is required")
	}

	if len(posts.Title) < 20 {
		return errors.New("title must be at least 20 characters")
	}

	if strings.TrimSpace(posts.Content) == "" {
		return errors.New("content is required")
	}

	if len(posts.Content) < 200 {
		return errors.New("content must be at least 200 characters")
	}

	if strings.TrimSpace(posts.Category) == "" {
		return errors.New("category is required")
	}

	if len(posts.Category) < 3 {
		return errors.New("category must be at least 3 characters")
	}

	if strings.TrimSpace(posts.Status) == "" {
		return errors.New("status is required")
	}

	if posts.Status != "publish" && posts.Status != "draft" && posts.Status != "thrash" {
		return errors.New("status must be 'publish', 'draft', or 'thrash'")
	}

	if _, err := s.postRepo.GetPostsById(posts.Id); err != nil {
		return errors.New("data not found")
	}

	return s.postRepo.Update(posts)
}

func (s *PostService) Delete(id int) error {
	if _, err := s.postRepo.GetPostsById(id); err != nil {
		return errors.New("data not found")
	}

	return s.postRepo.Delete(id)
}

func (s *PostService) UpdateStatusTrash(id int) error {
	if _, err := s.postRepo.GetPostsById(id); err != nil {
		return errors.New("data not found")
	}

	return s.postRepo.UpdateStatusTrash(id)
}

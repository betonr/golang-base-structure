package service

import (
	"context"
	"errors"
	"time"

	"github.com/betonr/golang-base-structure/internal/domain"

	"github.com/google/uuid"
)

type ArticleService struct {
	repo domain.ArticleRepository
}

type CreateArticleInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateArticleInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewArticleService(r domain.ArticleRepository) *ArticleService {
	return &ArticleService{repo: r}
}

func (s *ArticleService) CreateArticle(ctx context.Context, title, content, author string) (*domain.Article, error) {
	article := domain.Article{
		ID:        uuid.New().String(),
		Title:     title,
		Content:   content,
		Author:    author,
		CreatedAt: time.Now(),
	}
	if err := s.repo.Create(ctx, article); err != nil {
		return nil, err
	}
	return &article, nil
}

func (s *ArticleService) ListArticles(ctx context.Context) ([]domain.Article, error) {
	return s.repo.FindAll(ctx)
}

func (s *ArticleService) GetArticleByID(ctx context.Context, id string) (*domain.Article, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ArticleService) UpdateArticle(ctx context.Context, id string, title string, content string) error {
	article, err := s.repo.FindByID(ctx, id)
	if err != nil || article == nil {
		return errors.New("article not found")
	}

	article.Title = title
	article.Content = content
	return s.repo.Update(ctx, *article)
}

func (s *ArticleService) DeleteArticle(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

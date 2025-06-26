package domain

import (
	"context"
	"time"
)

type Article struct {
	ID        string
	Title     string
	Content   string
	Author    string
	CreatedAt time.Time
}

type ArticleRepository interface {
	Create(ctx context.Context, article Article) error
	FindAll(ctx context.Context) ([]Article, error)
	FindByID(ctx context.Context, id string) (*Article, error)
	Update(ctx context.Context, article Article) error
	Delete(ctx context.Context, id string) error
}

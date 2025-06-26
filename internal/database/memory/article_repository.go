package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/betonr/golang-base-structure/internal/domain"
)

type MemoryArticleRepository struct {
	mu       sync.RWMutex
	articles map[string]domain.Article
}

func NewMemoryArticleRepository() *MemoryArticleRepository {
	return &MemoryArticleRepository{
		articles: make(map[string]domain.Article),
	}
}

func (r *MemoryArticleRepository) Create(ctx context.Context, article domain.Article) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.articles[article.ID] = article
	return nil
}

func (r *MemoryArticleRepository) FindAll(ctx context.Context) ([]domain.Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []domain.Article
	for _, p := range r.articles {
		result = append(result, p)
	}
	return result, nil
}

func (r *MemoryArticleRepository) FindByID(ctx context.Context, id string) (*domain.Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if article, ok := r.articles[id]; ok {
		return &article, nil
	}
	return nil, errors.New("article not found")
}

func (r *MemoryArticleRepository) Update(ctx context.Context, article domain.Article) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.articles[article.ID]; !ok {
		return errors.New("article not found")
	}
	r.articles[article.ID] = article
	return nil
}

func (r *MemoryArticleRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.articles[id]; !ok {
		return errors.New("article not found")
	}
	delete(r.articles, id)
	return nil
}

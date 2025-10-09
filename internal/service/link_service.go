package service

import (
	"context"

	"github.com/pixisprod/URL-shortener/internal/cache"
	"github.com/pixisprod/URL-shortener/internal/repository"
	"github.com/pixisprod/URL-shortener/internal/util/hash"
)

type LinkService struct {
	r *repository.LinkRepository
	h hash.Generator
	c cache.Cacher[string]
}

func NewLinkService(
	r *repository.LinkRepository,
	h hash.Generator,
	c cache.Cacher[string],
) *LinkService {
	return &LinkService{r: r, h: h, c: c}
}

func (ls *LinkService) GenerateLink(url string) (string, error) {
	h, err := ls.h.Generate()
	if err != nil {
		return "", err
	}
	err = ls.r.Add(h, url)
	if err != nil {
		return "", err
	}
	return h, nil
}

func (ls *LinkService) GetLink(ctx context.Context, hash string) (string, error) {
	u, err := ls.c.Get(ctx, hash)
	if err == nil && u != "" {
		return u, nil
	}

	u, err = ls.r.GetByHash(hash)
	if err != nil {
		return "", err
	}
	_ = ls.c.Set(ctx, hash, u)

	return u, nil
}

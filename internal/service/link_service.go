package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/pixisprod/URL-shortener/internal/cache"
	"github.com/pixisprod/URL-shortener/internal/domain"
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

func (ls *LinkService) GenerateLink(url string, ll time.Time) (string, error) {
	h, err := ls.h.Generate()
	if err != nil {
		return "", err
	}
	link := domain.Link{
		Hash:      h,
		URL:       url,
		ExpiresAt: ll,
	}
	err = ls.r.Add(link)
	if err != nil {
		return "", err
	}
	return h, nil
}

func (ls *LinkService) GetLink(ctx context.Context, hash string) (string, error) {
	cachedURL, err := ls.c.Get(ctx, hash)
	if err == nil && cachedURL != "" {
		return cachedURL, nil
	}

	dl, err := ls.r.GetByHash(hash)
	if errors.Is(err, sql.ErrNoRows) {
		return "", domain.ErrLinkNotFound
	} else if err != nil {
		log.Println("Mystery error")
		return "", err
	}
	if time.Now().After(dl.ExpiresAt) {
		return "", domain.ErrLinkExpired
	}

	ttl := int(time.Until(dl.ExpiresAt).Seconds())
	if ttl > 0 {
		_ = ls.c.Set(ctx, hash, dl.URL, ttl)
	}

	return dl.URL, nil
}

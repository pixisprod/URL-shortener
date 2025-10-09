package service

import (
	"github.com/pixisprod/URL-shortener/internal/repository"
	"github.com/pixisprod/URL-shortener/internal/util/hash"
)

type LinkService struct {
	r *repository.LinkRepository
	h hash.Generator
}

func NewLinkService(r *repository.LinkRepository, h hash.Generator) *LinkService {
	return &LinkService{r: r, h: h}
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

func (ls *LinkService) GetLink(hash string) (string, error) {
	l, err := ls.r.GetByHash(hash)
	if err != nil {
		return "", err
	}
	return l, nil
}

package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pixisprod/url-shortener-link-service/internal/domain"
	"github.com/pixisprod/url-shortener-link-service/internal/model"
	"github.com/pixisprod/url-shortener-link-service/internal/service"
)

type LinkController struct {
	ls *service.LinkService
}

func NewLinkController(service *service.LinkService) *LinkController {
	return &LinkController{
		ls: service,
	}
}

func (lc *LinkController) GenLink(c *gin.Context) {
	s := model.ShortLink{}
	err := c.BindJSON(&s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid request body"})
		return
	}
	ttl := time.Now().Add(time.Duration(s.TTL) * time.Second)
	h, err := lc.ls.GenerateLink(s.Link, ttl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error while generating hash"})
		return
	}
	c.JSON(200, gin.H{"new_hash": h})
}

func (lc *LinkController) Redirect(c *gin.Context) {
	ctx := c.Request.Context()
	hash := c.Param("hash")
	url, err := lc.ls.GetLink(ctx, hash)
	if errors.Is(err, domain.ErrLinkNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	} else if errors.Is(err, domain.ErrLinkExpired) {
		c.JSON(http.StatusGone, gin.H{"msg": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Unknown error"})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pixisprod/URL-shortener/internal/model"
	"github.com/pixisprod/URL-shortener/internal/service"
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
	h, err := lc.ls.GenerateLink(s.Link)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error while generating hash"})
		return
	}
	c.JSON(200, gin.H{"new_hash": h})
}

func (lc *LinkController) Redirect(c *gin.Context) {
	hash := c.Param("hash")
	url, err := lc.ls.GetLink(hash)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Hash not found"})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

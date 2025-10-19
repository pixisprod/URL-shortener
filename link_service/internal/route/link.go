package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pixisprod/url-shortener-link-service/internal/controller"
)

func registerLinkRoutes(
	rg *gin.RouterGroup,
	c *controller.LinkController,
) {
	g := rg.Group("/links")
	g.POST("/cut", c.GenLink)
	g.GET("/r/:hash", c.Redirect)
}

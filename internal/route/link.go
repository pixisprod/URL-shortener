package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pixisprod/URL-shortener/internal/controller"
)

func registerLinkRoutes(
	rg *gin.RouterGroup,
	c *controller.LinkController,
) {
	g := rg.Group("/links")
	g.POST("/cut", c.GenLink)
	g.GET("/r/:hash", c.Redirect)
}

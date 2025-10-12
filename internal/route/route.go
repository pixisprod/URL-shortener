package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pixisprod/URL-shortener/internal/controller"
)

func RegisterRouters(
	e *gin.Engine,
	lc *controller.LinkController,
	sc *controller.ServiceController,
) {
	api := e.Group("/api")

	registerLinkRoutes(api, lc)
	registerServiceRoutes(api, sc)
}

func registerServiceRoutes(
	rg *gin.RouterGroup,
	c *controller.ServiceController,
) {
	g := rg.Group("/service")
	g.GET("/health", c.HealthCheck)
}

func registerLinkRoutes(
	rg *gin.RouterGroup,
	c *controller.LinkController,
) {
	g := rg.Group("/links")
	g.POST("/cut", c.GenLink)
	g.GET("/r/:hash", c.Redirect)
}

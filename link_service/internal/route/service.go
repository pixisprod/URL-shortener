package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pixisprod/url-shortener-link-service/internal/controller"
)

func registerServiceRoutes(
	rg *gin.RouterGroup,
	c *controller.ServiceController,
) {
	g := rg.Group("/service")
	g.GET("/health", c.HealthCheck)
}

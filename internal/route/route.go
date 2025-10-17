package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pixisprod/URL-shortener/internal/controller"
	"github.com/pixisprod/URL-shortener/internal/middleware"
)

func RegisterRouters(
	e *gin.Engine,
	lc *controller.LinkController,
	sc *controller.ServiceController,
	lm *middleware.LoggingMiddleware,
) {
	e.Use(lm.Log)
	api := e.Group("/api")

	registerLinkRoutes(api, lc)
	registerServiceRoutes(api, sc)
}

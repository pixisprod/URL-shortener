package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pixisprod/URL-shortener/internal/controller"
	"github.com/pixisprod/URL-shortener/internal/service"
)

func RegisterRouters(
	e *gin.Engine,
	ls *service.LinkService,
) {
	api := e.Group("/api")

	registerLinkRoutes(api, ls)
}

func registerLinkRoutes(
	rg *gin.RouterGroup,
	ls *service.LinkService,
) {
	g := rg.Group("/links")
	lc := controller.NewLinkController(ls)
	g.POST("/cut", lc.GenLink)
	g.GET("/r/:hash", lc.Redirect)
}

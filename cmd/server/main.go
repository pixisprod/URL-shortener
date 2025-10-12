package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pixisprod/URL-shortener/internal/cache"
	"github.com/pixisprod/URL-shortener/internal/config"
	"github.com/pixisprod/URL-shortener/internal/controller"
	"github.com/pixisprod/URL-shortener/internal/database"
	"github.com/pixisprod/URL-shortener/internal/repository"
	"github.com/pixisprod/URL-shortener/internal/route"
	"github.com/pixisprod/URL-shortener/internal/service"
	"github.com/pixisprod/URL-shortener/internal/util/hash"
)

func main() {
	cfg := config.LoadConfig()
	db, err := database.InitDbPool(
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SslMode,
		cfg.App.RetryInterval,
	)
	if err != nil {
		log.Println(err.Error())
	}
	rc := cache.NewRedisCacher(
		cache.InitRedisCacher(
			fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
			cfg.App.RetryInterval,
		),
	)
	h := hash.NewHashGenerator(cfg.Hash.Charset, cfg.Hash.Length)
	lr := repository.NewLinkRepository(db)
	ls := service.NewLinkService(lr, h, rc)
	lc := controller.NewLinkController(ls)

	sc := controller.NewServiceController()

	s := gin.Default()
	route.RegisterRouters(s, lc, sc)
	s.Run()
}

package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pixisprod/url-shortener-link-service/internal/cache"
	"github.com/pixisprod/url-shortener-link-service/internal/config"
	"github.com/pixisprod/url-shortener-link-service/internal/controller"
	"github.com/pixisprod/url-shortener-link-service/internal/database"
	"github.com/pixisprod/url-shortener-link-service/internal/middleware"
	"github.com/pixisprod/url-shortener-link-service/internal/repository"
	"github.com/pixisprod/url-shortener-link-service/internal/route"
	"github.com/pixisprod/url-shortener-link-service/internal/service"
	"github.com/pixisprod/url-shortener-link-service/internal/util/hash"
)

func main() {
	logger, err := setupLogger()
	if err != nil {
		log.Println(err.Error())
	}
	logger.Info("Server started!")
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

	lm := middleware.NewLoggingMiddleware(logger)
	s := gin.Default()
	route.RegisterRouters(s, lc, sc, lm)
	s.Run()
}

func setupLogger() (*slog.Logger, error) {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	h := slog.NewJSONHandler(file, &slog.HandlerOptions{
		AddSource: true,
	})
	return slog.New(h), nil
}

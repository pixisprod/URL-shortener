package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pixisprod/URL-shortener/internal/config"
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
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	h := hash.HashGenerator{
		Charset: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		Length:  5,
	}
	lr := repository.LinkRepository{Db: db}
	ls := service.NewLinkService(&lr, &h)

	e := gin.Default()
	route.RegisterRouters(e, ls)
	e.Run()
}

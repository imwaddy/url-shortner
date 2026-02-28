package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imwaddy/url-shortner/internal/config"
	"github.com/imwaddy/url-shortner/internal/db"
	"github.com/imwaddy/url-shortner/internal/handler"
	"github.com/imwaddy/url-shortner/internal/repository"
	"github.com/imwaddy/url-shortner/internal/service"
)

func main() {
	cfg := config.Load()

	database, err := db.NewMySQL(cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(database); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewURLRepository(database)
	svc := service.NewURLService(repo)
	h := handler.NewURLHandler(svc)

	r := gin.Default()
	h.RegisterRoutes(r)

	log.Fatal(r.Run(":" + cfg.Port))
}

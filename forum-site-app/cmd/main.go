package main

import (
	"forum-site-app/internal/configs"
	"forum-site-app/internal/handlers/memberships"
	"forum-site-app/internal/handlers/posts"
	membershipsRepo "forum-site-app/internal/repository/memberships"
	postsRepo "forum-site-app/internal/repository/posts"
	membershipsSvc "forum-site-app/internal/service/memberships"
	postsSvc "forum-site-app/internal/service/posts"
	"forum-site-app/pkg/internalsql"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var cfg *configs.Config

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile(
			"config",
		),
		configs.WithConfigType(
			"yaml",
		),
	)
	if err != nil {
		log.Fatal("Failed to initialize config: ", err)
	}

	cfg = configs.Get()
	log.Println("Config: ", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipsRepo := membershipsRepo.NewRepository(db)
	membershipsSvc := membershipsSvc.NewService(cfg, membershipsRepo)

	membershipsHandler := memberships.NewHandler(r, membershipsSvc)
	membershipsHandler.RegisterRoutes()

	postsRepo := postsRepo.NewRepository(db)
	postsSvc := postsSvc.NewService(cfg, postsRepo)

	postsHandler := posts.NewHandler(r, postsSvc)
	postsHandler.RegisterRoutes()

	r.Run(cfg.Service.Port)
}

package main

import (
	"log"
	"music-app/internal/configs"
	membershipsHandler "music-app/internal/handler/memberships"
	"music-app/internal/models/memberships"
	membershipsRepo "music-app/internal/repository/memberships"
	membershipsService "music-app/internal/service/memberships"
	"music-app/pkg/internalsql"

	"github.com/gin-gonic/gin"
)

func main() {
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
	db.AutoMigrate(&memberships.User{})

	r := gin.Default()

	membershipsRepo := membershipsRepo.NewRepository(db)
	membershipsService := membershipsService.NewService(cfg, membershipsRepo)
	membershipsHandler := membershipsHandler.NewHandler(r, membershipsService)
	membershipsHandler.RegisterRoutes()

	r.Run(cfg.Service.Port)
}

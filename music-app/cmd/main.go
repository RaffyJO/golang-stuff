package main

import (
	"log"
	"music-app/internal/configs"
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

	r := gin.Default()

	r.Run(cfg.Service.Port)
}

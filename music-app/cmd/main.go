package main

import (
	"log"
	"music-app/internal/configs"
	membershipsHandler "music-app/internal/handler/memberships"
	tracksHandler "music-app/internal/handler/tracks"
	"music-app/internal/models/memberships"
	"music-app/internal/models/track_activities"
	membershipsRepo "music-app/internal/repository/memberships"
	"music-app/internal/repository/spotify"
	tracksRepo "music-app/internal/repository/track_activities"
	membershipsService "music-app/internal/service/memberships"
	tracksService "music-app/internal/service/tracks"
	"music-app/pkg/httpclient"
	"music-app/pkg/internalsql"
	"net/http"

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
	db.AutoMigrate(&track_activities.TrackActivity{})

	r := gin.Default()

	httpclient := httpclient.NewClient(&http.Client{})
	spotifyOutbound := spotify.NewSpotifyOutbound(cfg, httpclient)
	tracksRepo := tracksRepo.NewRepository(db)
	tracksService := tracksService.NewService(spotifyOutbound, tracksRepo)
	tracksHandler := tracksHandler.NewHandler(r, tracksService)
	tracksHandler.RegisterRoutes()

	membershipsRepo := membershipsRepo.NewRepository(db)
	membershipsService := membershipsService.NewService(cfg, membershipsRepo)
	membershipsHandler := membershipsHandler.NewHandler(r, membershipsService)
	membershipsHandler.RegisterRoutes()

	r.Run(cfg.Service.Port)
}

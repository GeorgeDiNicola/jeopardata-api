package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"georgedinicola/jeopardy-api/api"
	"georgedinicola/jeopardy-api/internal/db"
)

func main() {
	db, err := db.CreateNewGormDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	jApi := &api.JeopardyApi{Db: db}

	router := gin.Default()

	router.GET("/contestants", jApi.GetAllContestants)
	router.GET("/episodes", jApi.GetEpisodes)
	router.GET("/episodes/:episodeNumber/performance", jApi.GetPerformanceForEpisodeNumber)
	router.GET("/games", jApi.GetAllGames)
	router.GET("/export", jApi.ExportAllGames)

	router.Run("localhost:8080")
}

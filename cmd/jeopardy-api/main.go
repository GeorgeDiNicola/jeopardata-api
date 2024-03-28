package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/georgedinicola/jeopardy-api/api"
	"github.com/georgedinicola/jeopardy-api/internal/db"
)

func main() {
	db, err := db.CreateNewDatabaseConnx()
	if err != nil {
		log.Fatal(err)
	}

	jApi := api.CreateNewJeopardyApi(db)

	router := gin.Default()

	router.GET("/contestants", jApi.GetAllContestants)
	router.GET("/games", jApi.GetAllGames)
	router.GET("/episodes", jApi.GetAllEpisodes)
	router.GET("/episodes/:episodeNumber/performance", jApi.GetPerformanceForEpisodeNumber)
	router.GET("/export", jApi.ExportAllGames)

	router.Run("localhost:8080")
}

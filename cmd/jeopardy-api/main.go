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

	router.GET("/episodes", jApi.GetEpisodes)

	router.Run("localhost:8080")
}

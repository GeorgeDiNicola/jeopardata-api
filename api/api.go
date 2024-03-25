package api

import (
	"georgedinicola/jeopardy-api/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JeopardyApi struct {
	Db db.DatabaseConnx
}

func (j *JeopardyApi) GetEpisodes(c *gin.Context) {
	allEpisodes, err := j.Db.GetAllEpisodes("DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, allEpisodes)
}

func (j *JeopardyApi) GetAllContestants(c *gin.Context) {
	allContestants, err := j.Db.GetAllContestants("ASC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, allContestants)
}

func (j *JeopardyApi) GetPerformanceForEpisodeNumber(c *gin.Context) {
	episodeNumber := c.Param("episodeNumber")

	episodes, err := j.Db.GetGameByEpisodeNumber(episodeNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, episodes)
}

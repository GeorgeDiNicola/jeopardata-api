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
	allEpisodes, err := j.Db.GetAllEpisodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, allEpisodes)
}

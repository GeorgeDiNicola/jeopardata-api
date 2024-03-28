package util

import (
	"math"

	"github.com/georgedinicola/jeopardy-api/internal/model"
	"github.com/gin-gonic/gin"
)

func CreatePaginationResponse(c *gin.Context, data interface{}, totalRecords, currentPage, limit int) {
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))
	var nextPage, prevPage *int

	if currentPage < totalPages {
		nextPageVal := currentPage + 1
		nextPage = &nextPageVal
	}

	if currentPage > 1 {
		prevPageVal := currentPage - 1
		prevPage = &prevPageVal
	}

	c.JSON(200, gin.H{
		"data": data,
		"pagination": model.PaginationMetaData{
			TotalRecords: totalRecords,
			CurrentPage:  currentPage,
			TotalPages:   totalPages,
			NextPage:     nextPage,
			PrevPage:     prevPage,
		},
	})
}

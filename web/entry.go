package web

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackparsonss/full-text-search/search"
)

var requestBody struct {
	Query string `json:"query" binding:"required"`
}

func Entry(index search.Index, documents []search.Document) {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/search", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		idxResult := index.Search(requestBody.Query)
		result := make([]string, len(idxResult))
		for i, v := range idxResult {
			result[i] = documents[v].Text
		}

		c.JSON(http.StatusOK, gin.H{"result": result})
	})
	r.Run()
}

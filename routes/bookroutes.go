package routes

import (
	"net/http"

	"bookcollection.com/rest-api/db"
	"github.com/gin-gonic/gin"
)

func getTradBooks(c *gin.Context) {
	books, err := db.GetAllTradBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error getting books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func getTradBooksByRef(c *gin.Context) {
	books, err := db.GetTradBooksByReference(c.Param("ref"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error getting books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

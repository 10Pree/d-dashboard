package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func listBooksHandler(c *gin.Context) {
	var books []Book

	if result := db.Find(&books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &books)
}

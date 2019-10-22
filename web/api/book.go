package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lancatlin/library/pkg/search"
)

func SearchBook(c *gin.Context) {
	keyword := c.Query("q")
	books := search.SearchBooks(keyword)
	c.JSON(200, books)
}

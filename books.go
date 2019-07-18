package main

import (
	"github.com/gin-gonic/gin"
)

func booksIndex(c *gin.Context) {
	book := Book{
		BookName: "自然文學之書",
		Authors: []Author{
			Author{Name: "林宏信"},
		},
		Publisher: Publisher{
			Name: "遠流出版社",
		},
		Year: 2019,
		Category: Category{
			Name: "生態美學",
		},
		ClassificationNumber: "121",
		Items: []Item{
			Item{
				Barcode: "A380",
				Status:  StatusInside,
			},
			Item{
				Barcode: "A381",
				Status:  StatusLending,
			},
		},
	}
	page := struct {
		User
		Book
	}{getUser(c), book}
	c.HTML(200, "books_index.html", page)
}

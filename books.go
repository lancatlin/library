package main

import (
	"github.com/gin-gonic/gin"
)

var fakeBooks = []Book{}
var fakeItems = []Item{}

func init() {
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
	fakeBooks = append(fakeBooks, book)
	fakeItems = []Item{
		Item{
			Barcode: "A380",
			Status:  StatusInside,
			Book:    fakeBooks[0],
		},
		Item{
			Barcode: "A381",
			Status:  StatusLending,
			Book:    fakeBooks[0],
		},
	}
}

func booksIndex(c *gin.Context) {
	page := struct {
		User
		Book
	}{getUser(c), fakeBooks[0]}
	c.HTML(200, "books_index.html", page)
}

func books(c *gin.Context) {
	page := struct {
		User
		Items []Item
	}{
		getUser(c),
		fakeItems,
	}
	c.HTML(200, "books.html", page)
}

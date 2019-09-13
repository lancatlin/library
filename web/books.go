package web

import (
	"github.com/gin-gonic/gin"
	"github.com/lancatlin/library/pkg/model"
)

var fakeBooks = []model.Book{}
var fakeItems = []model.Item{}

func init() {
	book := model.Book{
		Name: "自然文學之書",
		Authors: []model.Author{
			model.Author{Name: "林宏信"},
		},
		Publisher: model.Publisher{
			Name: "遠流出版社",
		},
		Year: 2019,
		Category: model.Category{
			Name: "生態美學",
		},
		ClassificationNumber: "121",
		Items: []model.Item{
			model.Item{
				Barcode: "A380",
			},
			model.Item{
				Barcode: "A381",
			},
		},
	}
	fakeBooks = append(fakeBooks, book)
	fakeItems = []model.Item{
		model.Item{
			Barcode: "A380",
			Book:    fakeBooks[0],
		},
		model.Item{
			Barcode: "A381",
			Book:    fakeBooks[0],
		},
	}
}

func booksIndex(c *gin.Context) {
	page := struct {
		model.User
		model.Book
	}{getUser(c), fakeBooks[0]}
	c.HTML(200, "books_index.html", page)
}

func books(c *gin.Context) {
	page := struct {
		model.User
		Items []model.Item
	}{
		getUser(c),
		fakeItems,
	}
	c.HTML(200, "books.html", page)
}

func booksNew(c *gin.Context) {
	page := struct {
		model.User
		Categories []model.Category
	}{
		getUser(c),
		[]model.Category{
			model.Category{Name: "自然文學"},
			model.Category{Name: "自然美學"},
			model.Category{Name: "自然生態"},
			model.Category{Name: "自然哲學"},
		},
	}
	c.HTML(200, "books_new.html", page)
}

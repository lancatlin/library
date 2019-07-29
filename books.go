package main

import (
	"log"
	"strconv"

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
			},
			Item{
				Barcode: "A381",
			},
		},
	}
	fakeBooks = append(fakeBooks, book)
	fakeItems = []Item{
		Item{
			Barcode: "A380",
			Book:    fakeBooks[0],
		},
		Item{
			Barcode: "A381",
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

func booksNew(c *gin.Context) {
	page := struct {
		User
		Categories []Category
	}{
		getUser(c),
		[]Category{
			Category{Name: "自然文學"},
			Category{Name: "自然美學"},
			Category{Name: "自然生態"},
			Category{Name: "自然哲學"},
		},
	}
	c.HTML(200, "books_new.html", page)
}

func (b *Book) newItem() Item {
	item := Item{
		Barcode:      b.genBarcode(),
		NewBookLabel: b.genNewLabel(),
		Book:         *b,
	}
	if err := db.Create(&item).Error; err != nil {
		log.Fatalln(err)
	}
	if err := db.Commit().Error; err != nil {
		log.Fatalln(err)
	}
	return item
}

func (b Book) genBarcode() string {
	var count []int
	query := `
	SELECT COUNT(1) FROM books b
	INNER JOIN items i ON b.id = i.book_id
	WHERE b.category_id = ?
	`
	if err := db.Raw(query, b.Category.ID).Scan(&count).Error; err != nil {
		log.Fatalln(err)
	}
	return b.Category.Prefix + strconv.Itoa(count[0]+1)
}

func (b Book) genNewLabel() string {
	var count int
	if err := db.Where(&Item{NewBookLabel: ""}).Find(&[]Item{}).Count(&count).Error; err != nil {
		log.Fatalln(err)
	}
	return "N" + strconv.Itoa(count+1)
}

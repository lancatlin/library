package main

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
)

func booksImport(file io.Reader) (err error) {
	r := csv.NewReader(file)
	r.Comma = ' '
	// 第一行不讀
	columes, err := r.Read()
	if err != nil {
		return err
	}
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		data := make(map[string]string)
		for i, v := range columes {
			data[v] = line[i]
		}
		log.Println(data)

		book := Book{
			BookName:             data["BookName"],
			ISBN:                 data["ISBN"],
			Description:          data["Description"],
			Cover:                data["CoverImage"],
			ClassificationNumber: data["ClassificationNumber"],
		}
		authors := strings.Split(data["Authors"], ",")
		book.Authors = make([]Author, len(authors))
		// 找尋已經有的作家，如果存在就使用，否則創建
		for i, v := range authors {
			var author Author
			res := db.FirstOrInit(&author, Author{Name: v})
			if res.Error != nil {
				log.Fatalln(res.Error)
			}
			log.Println(author)
			book.Authors[i] = author
		}
		// 出版社
		var publisher Publisher
		if err = db.FirstOrInit(&publisher, Publisher{Name: data["Publisher"]}).Error; err != nil {
			log.Fatalln(err)
		}
		book.Publisher = publisher
		// 分類
		var category Category
		if err = db.Where("name = ?", data["Category"]).First(&category).Error; err != nil {
			log.Fatalln(err)
		}
		log.Println(category)
		book.Category = category
		// 標籤
		tags := strings.Split(data["Tags"], ",")
		log.Println(tags)
		book.Tags = make([]Tag, len(tags))
		for i, v := range tags {
			var tag Tag
			if err = db.FirstOrCreate(&tag, Tag{Name: v}).Error; err != nil {
				log.Fatalln(err)
			}
			log.Println(tag)
			book.Tags[i] = tag
		}
		// 館藏
		amount, err := strconv.Atoi(data["ItemsAmount"])
		if err != nil {
			log.Println(err)
		}
		log.Println(amount)
		book.Items = make([]Item, amount)
		for i := 0; i < amount; i++ {
			item := book.newItem()
			log.Println(item)
		}
		// 年代
		year, err := strconv.Atoi(data["Year"])
		if err != nil {
			log.Fatalln(err)
		}
		book.Year = year
		// Create
		if err = db.Create(&book).Error; err != nil {
			log.Fatal(err)
		}
		db.First(&book, book.ID)
		log.Println(book)
	}
	db.Commit()
	return nil
}

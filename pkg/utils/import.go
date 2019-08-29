package utils

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/lancatlin/library/pkg/model"
)

var (
	// ErrCategoryNotDefined mean the ID contain an undefined category prefix
	ErrCategoryNotDefined = errors.New("utils: category is undefined")
	// ErrInvalidID mean the ID is invalid
	ErrInvalidID = errors.New("utils: ID is invalid")
)

// ImportBooks load a csv file an parse it into database
// the csv file need the columns below:
// ID, BookName, Authors, Publisher, ISBN, Description, CoverImage, ClassificationNumber, Year, Tags
// ID must be an defined category prefix and a number, like A147, C5692... in regex: '^[A-Z][0-9]+$'
// Authors must be a string split with ',' or ';'. the content will be split by '([,;、，\n] *)+'
// Tags is a string join by ',': also split by '([,;、，\n] *)+'
func ImportBooks(file io.Reader) (err error) {
	r := csv.NewReader(file)
	r.Comma = ' '
	// 第一行不讀
	columns, err := r.Read()
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
		for i, v := range columns {
			data[v] = line[i]
		}
		log.Println(data)
		if err := parse(data); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func parse(data map[string]string) (err error) {
	book := model.Book{
		Name:                 data["BookName"],
		ISBN:                 data["ISBN"],
		Description:          data["Description"],
		Cover:                data["CoverImage"],
		ClassificationNumber: data["ClassificationNumber"],
	}
	book.Authors = parseAuthors(data["Authors"])
	// 出版社
	var publisher model.Publisher
	if err = db.FirstOrInit(&publisher, model.Publisher{Name: data["Publisher"]}).Error; err != nil {
		log.Fatalln(err)
	}
	book.Publisher = publisher
	// 標籤
	tags := strings.Split(data["Tags"], ",")
	log.Println(tags)
	book.Tags = make([]model.Tag, len(tags))
	for i, v := range tags {
		var tag model.Tag
		if err = db.FirstOrCreate(&tag, model.Tag{Name: v}).Error; err != nil {
			log.Fatalln(err)
		}
		log.Println(tag)
		book.Tags[i] = tag
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
	return nil
}

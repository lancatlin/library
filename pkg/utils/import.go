package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"time"

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
// ID, BookName, Authors, Publisher, ISBN, Description, CoverImage, ClassNum, Year, Tags
// ID must be an defined category prefix and a number, like A147, C5692... in regex: '^[A-Z][0-9]+$'
// Authors must be a string split with ',' or ';'. the content will be split by '([,;、，\n] *)+'
// Tags is a string join by ',': also split by '([,;、，\n] *)+'
// If a book has multiple items, write the ID only, leave the others blank.
func ImportBooks(file io.Reader) (errorMessages []string) {
	raws := readRaws(file)
	errChan := make(chan string, 5)
	for _, raw := range raws {
		go importBook(raw, errChan)
	}
	count := 0
	for {
		select {
		case msg := <-errChan:
			if msg != "" {
				errorMessages = append(errorMessages, msg)
			}
			count++
			if count >= len(raws) {
				return
			}
		case <-time.After(time.Second * 2 * time.Duration(len(raws))):
			return
		}
	}
}

func readRaws(file io.Reader) (raws []map[string]string) {
	r := csv.NewReader(file)
	allRaws, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	if len(allRaws) <= 1 {
		return
	}
	columns := allRaws[0]
	raws = make([]map[string]string, len(allRaws)-1)
	for i, line := range allRaws[1:] {
		raws[i] = convertToMap(columns, line)
	}
	return
}

func convertToMap(columns, src []string) (dest map[string]string) {
	dest = make(map[string]string)
	for i, column := range columns {
		dest[column] = src[i]
	}
	return
}

func importBook(data map[string]string, errChan chan string) {
	book := model.Book{
		Name:        data["BookName"],
		Description: data["Description"],
		Cover:       data["CoverImage"],
	}
	book.Authors = parseAuthors(data["Authors"])
	book.ISBN = parseISBN(data["ISBN"])
	// 出版社
	book.Publisher = parsePublisher(data["Publisher"])
	// 標籤
	book.Tags = parseTagsAndCreate(data["Tags"])
	// 年代
	book.Year = parseYear(data["Year"])
	book.ClassNums = parseClassNum(data["ClassNum"])
	supporters := parseSupporter(data["Supporter"])
	barcodes := parseBarcodes(data["Barcodes"])
	if len(barcodes) == 0 {
		errChan <- fmt.Sprintf("Error: %s Cannot create item without barcodes", book.Name)
		return
	}
	if len(supporters) > len(barcodes) {
		errChan <- fmt.Sprintf("Error: %s's supporters are more than items", book.Name)
		return
	}
	var err error
	book.Category, err = model.GetCategory(barcodes)
	if err != nil {
		errChan <- fmt.Sprintf("Error: %s %s", book.Name, err.Error())
		return
	}
	err = book.InitItems(barcodes, supporters)
	if err != nil {
		errChan <- err.Error()
	}
	// Create
	if err = db.Create(&book).Error; err != nil {
		errChan <- fmt.Sprintf("%s create error: %s", book.Name, err.Error())
	}
	errChan <- ""
}

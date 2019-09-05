package utils

import (
	"encoding/csv"
	"errors"
	"io"
	"log"

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
func ImportBooks(file io.Reader) (errList []error) {
	raws := readRaws(file)
	for _, raw := range raws {
		err := parse(raw)
		if err != nil {
			errList = append(errList, err)
		}
	}
	return
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

func parse(data map[string]string) (err error) {
	book := model.Book{
		Name:        data["BookName"],
		Description: data["Description"],
		Cover:       data["CoverImage"],
	}
	book.Authors = parseAuthors(data["Authors"])
	book.ISBN, err = parseISBN(data["ISBN"])
	if err != nil {
		return
	}
	// 出版社
	book.Publisher = parsePublisher(data["Publisher"])
	// 標籤
	book.Tags = parseTagsAndCreate(data["Tags"])
	// 年代
	book.Year = parseYear(data["Year"])
	// Create
	if err = db.Create(&book).Error; err != nil {
		log.Fatal(err)
	}
	supporters := parseSupporter(data["Supporter"])
	barcodes := parseBarcodes(data["Barcodes"])
	if len(supporters) > len(barcodes) {
		err = errors.New("Invalid supporter number")
		return
	}
	book.Items = make([]model.Item, len(barcodes))
	for i, barcode := range barcodes {
		supporter := ""
		if len(supporters)-1 >= i {
			supporter = supporters[i]
		}
		var item model.Item
		item, err = book.NewItemWithBarcode(barcode, supporter)
		if err != nil {
			log.Println(err)
			return
		}
		book.Items[i] = item
	}
	db.Save(&book)
	log.Println(book)
	return nil
}

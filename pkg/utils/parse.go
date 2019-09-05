package utils

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/jinzhu/gorm"
	"github.com/lancatlin/library/pkg/model"
)

var (
	// ErrInvalidISBNLength is threw when isbn is in a invalid length(not 10 or 13)
	ErrInvalidISBNLength = errors.New("utils: the ISBN length is invalid")
	// ErrISBNParseError is threw when strconv.Atoi throw an error
	ErrISBNParseError = errors.New("utils: the ISBN is invalid")
)

func parseAuthors(s string) (result []model.Author) {
	splitChar := "||"
	newS, err := regexp2.MustCompile(`[,;、，\n] ?(?![^\(]*\))`, regexp2.RE2).Replace(s, splitChar, -1, -1)
	authors := strings.Split(newS, splitChar)
	if err != nil {
		panic(err)
	}
	result = make([]model.Author, len(authors))
	// 找尋已經有的作家，如果存在就使用，否則創建
	for i, v := range authors {
		var author model.Author
		res := db.FirstOrInit(&author, model.Author{Name: string(v)})
		if res.Error != nil {
			log.Fatalln(res.Error)
		}
		result[i] = author
	}
	return
}

func parseCategory(s string) (c model.Category, err error) {
	if !regexp.MustCompile("^[A-Z][a-z]*[0-9]+$").MatchString(s) {
		err = ErrInvalidID
		return
	}
	prefix := regexp.MustCompile("^[A-Z][a-z]*").FindString(s)
	if err = db.Where("prefix = ?", prefix).First(&c).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = ErrCategoryNotDefined
			return
		}
		panic(err)
	}
	return
}

// parseYear parse a string to int
// if the year isn't Common Era, considered ROC
func parseYear(s string) (year int) {
	yearString := regexp.MustCompile(`\d+`).FindString(s)
	year, err := strconv.Atoi(yearString)
	if err != nil {
		log.Println(err)
		return 0
	}
	if len(yearString) != 4 {
		// Not C.E.
		// Treat as ROC
		year += 1911
	}
	return
}

func parseISBN(s string) (isbn int, err error) {
	isbnString := regexp.MustCompile(`[- \s]`).ReplaceAllString(s, "")
	if l := len(isbnString); l != 10 && l != 13 {
		err = ErrInvalidISBNLength
		return
	}
	isbn, err = strconv.Atoi(isbnString)
	if err != nil {
		err = ErrISBNParseError
		return
	}
	return
}

func parseClassNum(s string) (nums []model.ClassNum) {
	matches := regexp.MustCompile(`\d+\.?\d*`).FindAllString(s, -1)
	nums = make([]model.ClassNum, len(matches))
	for i, v := range matches {
		var err error
		nums[i], err = model.NewClassNum(v)
		if err != nil {
			panic(err)
		}
	}
	return
}

func parseBarcodes(s string) []string {
	return regexp.MustCompile(`[A-Z][a-z]*\d+`).FindAllString(s, -1)
}

func parsePublisher(s string) (publisher model.Publisher) {
	if err := db.FirstOrInit(&publisher, model.Publisher{Name: s}).Error; err != nil {
		panic(err)
	}
	return
}

func splitByCommaAndSemicolon(s string) []string {
	return regexp.MustCompile(`[,;] *`).Split(s, -1)
}

func parseTagsAndCreate(s string) (tags []model.Tag) {
	tagsName := splitByCommaAndSemicolon(s)
	tags = make([]model.Tag, len(tagsName))
	for i, name := range tagsName {
		tags[i] = findOrCreateTag(name)
	}
	return
}

func findOrCreateTag(name string) (tag model.Tag) {
	if err := db.FirstOrCreate(&tag, model.Tag{Name: name}).Error; err != nil {
		panic(err)
	}
	return
}

func parseSupporter(s string) (supporters []string) {
	return splitByCommaAndSemicolon(s)
}

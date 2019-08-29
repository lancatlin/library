package utils

import (
	"log"
	"regexp"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/jinzhu/gorm"
	"github.com/lancatlin/library/pkg/model"
)

func parseAuthors(s string) (result []model.Author) {
	splitChar := "||"
	newS, err := regexp2.MustCompile(`[,;、，\n] ?(?![^\(]*\))`, regexp2.RE2).Replace(s, splitChar, -1, -1)
	authors := strings.Split(newS, splitChar)
	if err != nil {
		panic(err)
	}
	log.Println(authors)
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

func getCategory(s string) (c model.Category, err error) {
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

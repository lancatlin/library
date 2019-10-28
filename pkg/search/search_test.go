package search

import (
	"sync"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/lancatlin/library/pkg/account"
	"github.com/lancatlin/library/pkg/model"
)

var db *gorm.DB

var searcher searchImpl

func init() {
	filename := "search.sqlite"
	var err error
	db, err = gorm.Open("sqlite3", filename)
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&model.Book{}, &model.Item{}, &model.Account{},
		&model.Record{}, &model.Category{}, &model.Publisher{}, &model.Author{},
		&model.Tag{}).Error; err != nil {
		panic(err)
	}
	model.SetDB(db)
	model.InitCategoriesFromConfigs()
	err = account.LoadAllAccounts(db)
	if err != nil {
		panic(err)
	}

	searcher = searchImpl{
		db:   db,
		lock: &sync.Mutex{},
	}
}

func TestSearchBook(t *testing.T) {
	word := "台灣"
	column := "name"
	answer := []string{
		"玉峯觀止-台灣自然、宗教與教育之我見",
		"台灣必須廢核的10個理由",
		"自己的台灣自己救",
		"21世紀台灣主流的土石亂流",
		"台灣蝴蝶-食草與蜜源植物大圖鑑",
		"山災地變人造孽-21世紀台灣主流的土石亂流",
		"我的水中夥伴-生物學家談台灣溪流魚類和環境故事",
	}
	checkEqualBooks(word, column, answer, t)

	word = "957"
	column = "isbn"
	answer = []string{
		"共享自然的喜悅",
		"與孩子分享自然",
		"玉峯觀止-台灣自然、宗教與教育之我見",
		"傾聽自然",
		"天空的眼睛",
		"餐桌上的家鄉",
		"蘇建和案21年生死簿:蘇友辰律師口述史",
		"我生命中的花草樹木",
		"世界上最古老的生物",
		"敏督利注",
		"21世紀台灣主流的土石亂流",
		"逃/我們的寶島，他們的牢",
		"泰雅傳統竹屋",
		"織起一座彩虹祖靈橋",
		"阿里山物語(下)",
		"夢想南極 : 荒冰野地的魅力",
		"山災地變人造孽-21世紀台灣主流的土石亂流",
		"臺灣素人 : 宗教、精神、價值與人格",
		"女農討山誌-一個女子與土地的深情記事",
		"雪山綠緣-雪山地區植物選介",
		"捍衛正義 : 烏山頭水庫保衛戰",
		"展讀大坑天書",
		"綠島金夢",
		"臺灣生態社區的故事",
	}
	checkEqualBooks(word, column, answer, t)
}

func checkEqualBooks(word, column string, answer []string, t *testing.T) {
	var books []model.Book
	if err := searcher.searchByColumn(word, &books, column); err != nil {
		t.Error(err)
	}
	if len(books) != len(answer) {
		t.Errorf("answer length not equal: %v %v", answer, books)
	}
	for i, book := range books {
		if book.Name != answer[i] {
			t.Errorf(`Answer not equal: want %s have %s`, answer[i], book.Name)
		}
	}
}

/*
Not using
type dog int

func (d dog) Equal(obj interface{}) bool {
	if theDog, ok := obj.(dog); ok {
		if int(theDog) == int(d) {
			return true
		}
	}
	return false
}

func TestMerge(t *testing.T) {
	d1 := []model.Merger{dog(1), dog(2), dog(3), dog(5), dog(6)}
	d2 := []model.Merger{dog(3), dog(4), dog(5), dog(8), dog(9)}
	answer := []int{1, 2, 3, 5, 6, 4, 8, 9}
	var result []model.Merger
	merge(d1, d2, result)
	for i := range answer {
		dog := result[i].(dog)
		if int(dog) != answer[i] {
			t.Errorf("Answer not equal: want %d got %d", answer[i], dog)
		}
	}
}
*/

func TestSearchAccounts(t *testing.T) {
	keywords := []string{
		"林",
		"陳",
		"志",
		"78",
		"365",
		"574",
		``,
	}
	for _, keyword := range keywords {
		answer := searcher.fakeSearchAccount(keyword)
		result := searcher.SearchAccounts(keyword)
		compare(answer, result, t)
	}
}

func (s searchImpl) fakeSearchAccount(keyword string) (accounts []model.Account) {
	err := s.db.Where("name LIKE ?", "%"+keyword+"%").Or("phone LIKE ?", "%"+keyword+"%").Find(&accounts).Error
	if err != nil {
		panic(err)
	}
	return
}

func compare(answer, result []model.Account, t *testing.T) {
	if len(answer) != len(result) {
		t.Errorf(`Not eqaul: result: %v\nanswer: %v`, result, answer)
	}
	for i := range result {
		if result[i].ID != answer[i].ID {
			t.Errorf(`Not eqaul in %d: result: %v\nanswer: %v`, i, result[i], answer[i])
		}
	}
}

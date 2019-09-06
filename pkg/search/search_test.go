package search

import (
	"testing"

	"github.com/lancatlin/library/pkg/model"
)

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
	if err := searchByColumn(&books, word, column); err != nil {
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
	result := merge(d1, d2)
	for i := range answer {
		dog := result[i].(dog)
		if int(dog) != answer[i] {
			t.Errorf("Answer not equal: want %d got %d", answer[i], dog)
		}
	}
}

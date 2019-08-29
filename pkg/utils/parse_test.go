package utils

import "testing"

func TestParseAuthors(t *testing.T) {
	data := map[string][]string{
		"戴立克.簡申(Derrick Jensen), 喬治.德芮芬(George Draffan)": []string{"戴立克.簡申(Derrick Jensen)", "喬治.德芮芬(George Draffan)"},
		"林志欽、卓子右、姚明堯、劉晉榮、郭信利、蘇伯瀚、謝億文、朱珉寬":                []string{"林志欽", "卓子右", "姚明堯", "劉晉榮", "郭信利", "蘇伯瀚", "謝億文", "朱珉寬"},
		"陳維滄": []string{"陳維滄"},
		"亞歷塞維奇 (Aleksievich, Svetlana)": []string{"亞歷塞維奇 (Aleksievich, Svetlana)"},
		"薩斯曼 (Sussman, Rachel)":         []string{"薩斯曼 (Sussman, Rachel)"},
		"何貞青，鄭一青，林琮盛":                   []string{"何貞青", "鄭一青", "林琮盛"},
		"林為道(Baunay.Watan)":             []string{"林為道(Baunay.Watan)"},
	}
	for q, a := range data {
		authors := parseAuthors(q)
		if len(authors) != len(a) {
			t.Errorf("Answer not equal: want %s have %s", a, authors)
		}
		for i := range authors {
			if a[i] != authors[i].String() {
				t.Errorf("Answer not equal: want %s have %s\n", a, authors)
			}
		}
	}
}
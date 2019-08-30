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

func TestParseYear(t *testing.T) {
	data := map[string]int{
		"2011年":   2011,
		"94年":     2005,
		"1994年二版": 1994,
		"106":     2017,
		"2018":    2018,
		" ":       0,
		"第2018年":  2018,
		"107.3":   2018,
		"一百零七年":   0,
	}
	for q, a := range data {
		if y := parseYear(q); y != a {
			t.Errorf("Year not equal: want %d have %d\n", a, y)
		}
	}
}

func TestParseISBN(t *testing.T) {
	data := map[string]int{
		"978-986-89294-9-4": 9789868929494,
		"978-957-13-6810-8": 9789571368108,
		"9789869611732":     9789869611732,
		"979 986 961 1732":  9799869611732,
		"9799869611732\n":   9799869611732,
	}
	wrong := map[string]error{
		"978986892949":     ErrInvalidISBNLength,
		"9a8957-13-6810-8": ErrISBNParseError,
		"979 986 961 1732": nil,
	}
	for q, a := range data {
		if r, err := parseISBN(q); r != a {
			if err != nil {
				t.Error(err)
			}
			t.Errorf("ISBN not equal: want %d have %d\n", a, r)
		}
	}
	for q, a := range wrong {
		if _, err := parseISBN(q); err != a {
			t.Errorf("ISBN error not equal: want %s have %s", a, err)
		}
	}
}

func TestParseClassNum(t *testing.T) {
	data := map[string][]string{
		"783.3886":          []string{"783.3886"},
		"733.9/121.9/129.4": []string{"733.9", "121.9", "129.4"},
		"486":               []string{"486"},
		"796 758.21":        []string{"796", "758.21"},
	}
	for q, a := range data {
		c := parseClassNum(q)
		if len(c) != len(a) {
			t.Errorf("ClassNum not equal: want %s have %s", a, c)
		}
		for i := range c {
			if a[i] != c[i].String() {
				t.Errorf("ClassNum not equal: want %s have %s", a, c)
			}
		}
	}
}

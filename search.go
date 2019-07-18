package main

import "github.com/gin-gonic/gin"

func search(c *gin.Context) {
	page := struct {
		User
		Results []Book
		Query   string
	}{
		getUser(c),
		[]Book{
			Book{
				BookName: "能高越嶺道——穿越時空之旅",
				Authors: []Author{
					Author{
						Name: "楊南郡",
					},
					Author{
						Name: "徐如林",
					},
				},
				Publisher: Publisher{Name: "測試出版社"},
				Cover:     "https://im1.book.com.tw/image/getImage?i=https://www.books.com.tw/img/001/074/16/0010741610.jpg&v=58775ac6&w=348&h=348",
			},
			Book{
				BookName: "能高越嶺道——穿越時空之旅2",
				Authors: []Author{
					Author{
						Name: "作者3",
					},
				},
				Publisher: Publisher{Name: "測試出版社"},
				Cover:     "https://im1.book.com.tw/image/getImage?i=https://www.books.com.tw/img/001/074/16/0010741610.jpg&v=58775ac6&w=348&h=348",
			},
		},
		"能高",
	}
	c.HTML(200, "search_results.html", page)
}

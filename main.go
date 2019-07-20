package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func loadTemplate() (tpl *template.Template) {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	return
}

func main() {
	r := gin.Default()
	r.SetHTMLTemplate(loadTemplate())
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", getUser(c))
	})
	r.GET("/books/index", booksIndex)
	r.GET("/search/simple", func(c *gin.Context) {
		c.HTML(200, "search_simple.html", getUser(c))
	})
	r.GET("/search/detailed", func(c *gin.Context) {
		page := struct {
			User
			Categories []Category
		}{
			getUser(c),
			[]Category{
				Category{Name: "自然文學"},
				Category{Name: "自然美學"},
				Category{Name: "自然生態"},
				Category{Name: "自然哲學"},
			},
		}
		c.HTML(200, "search_detailed.html", page)
	})
	r.GET("/search", search)
	r.GET("/lending", func(c *gin.Context) {
		c.HTML(200, "lending.html", getUser(c))
	})
	r.GET("/return", func(c *gin.Context) {
		c.HTML(200, "return.html", getUser(c))
	})
	r.Run(":8080")
}

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
	r.Run(":8080")
}

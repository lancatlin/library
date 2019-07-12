package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func loadTemplate() (tpl *template.Template) {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	return
}

func main() {
	r := gin.New()
	r.SetHTMLTemplate(loadTemplate())
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.gohtml", nil)
	})
	r.Run(":8080")
}

package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lancatlin/library/pkg/model"
)

func Register(r *gin.Engine) {
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
			model.User
			Categories []model.Category
		}{
			getUser(c),
			[]model.Category{
				model.Category{Name: "自然文學"},
				model.Category{Name: "自然美學"},
				model.Category{Name: "自然生態"},
				model.Category{Name: "自然哲學"},
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
	r.GET("/management/books", books)
	r.GET("/management/books/new", booksNew)
	r.GET("/management/books/import", func(c *gin.Context) {
		user := getUser(c)
		if user.Role != model.RoleAdmin {
			c.HTML(401, "401.html", user)
		}
		c.HTML(200, "books_import.html", user)
	})
	r.POST("/management/books/import", func(c *gin.Context) {
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.Error(err)
			return
		}
		err = booksImport(file)
		if err != nil {
			log.Println(err)
			c.Error(err)
			return
		}
		c.String(200, "%s", "匯入成功")
	})
}

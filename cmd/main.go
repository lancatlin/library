package main

import (
	"../web"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob()
	router.Register(r)
	r.Run("localhost:8080")
}

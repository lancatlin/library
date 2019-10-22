package handler

import (
	"../pkg/model"
	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) model.User {
	return model.GetUser(c.Cookie("session"))
}

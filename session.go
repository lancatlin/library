package main

import "github.com/gin-gonic/gin"

func getUser(c *gin.Context) User {
	return User{
		UserName: "TestUser",
		Role:     RoleAdmin,
		Login:    true,
	}
}

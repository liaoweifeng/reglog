package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var account map[string]User = make(map[string]User)

func main() {
	router := gin.Default()
	//login
	router.POST("/login", func(c *gin.Context) {
		var user User

		err := c.Bind(&user)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		user.Password = c.PostForm("password")
		user.Username = c.PostForm("username")
		if v, ok := account[user.Username]; ok && v.Password == user.Password {
			c.JSON(http.StatusOK, gin.H{
				"username": v.Username,
				"password": v.Password,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "账号或者密码有误",
			})
		}
	})

	//register
	router.POST("/register", func(c *gin.Context) {
		var user User
		err := c.Bind(&user)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		user.Username = c.DefaultPostForm("username", "guest")
		username := user.Username
		user.Password = c.PostForm("password")
		if _, ok := account[username]; ok {
			message := "用户名" + username + "已存在"
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": message,
			})
		} else {
			account[username] = user
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "注册成功",
			})
		}
	})
	router.Run(":8080")
}

type User struct {
	Username string
	Password string
}

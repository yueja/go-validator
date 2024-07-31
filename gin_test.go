package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type Register struct {
	User     string `form:"user" json:"user" validate:"required,min=3,max=10" zhtrans:"用户名"`
	Age      int    `form:"age" json:"age" validate:"required,gte=0,lte=120" zhtrans:"年龄"`
	Sex      string `form:"sex" json:"sex" validate:"required,oneof=woman man" zhtrans:"性别"`
	Password string `form:"password" json:"password" validate:"required,min=8,max=20" zhtrans:"密码"`
}

func TestGin(t *testing.T) {
	router := gin.Default()

	// 绑定JSON的示例 ({"user": "wade", "password": "999999"})
	router.POST("/registerJSON", func(c *gin.Context) {
		var register Register
		if err := ShouldBindGinValidator(c, &register); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("register info:%#v\n", register)
		c.JSON(http.StatusOK, gin.H{
			"user":     register.User,
			"password": register.Password,
		})
	})

	// 绑定form表单示例 (user=wade&password=999999)
	router.POST("/registerForm", func(c *gin.Context) {
		var register Register
		if err := ShouldBindGinValidator(c, &register); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("register info:%#v\n", register)
		c.JSON(http.StatusOK, gin.H{
			"user":     register.User,
			"password": register.Password,
		})
	})

	// 绑定QueryString示例 (/registerForm?user=wade&password=999999)
	router.GET("/registerForm", func(c *gin.Context) {
		var register Register
		if err := ShouldBindGinValidator(c, &register); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("register info:%#v\n", register)
		c.JSON(http.StatusOK, gin.H{
			"user":     register.User,
			"password": register.Password,
		})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

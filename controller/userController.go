package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"wzjwh/dao"
	"wzjwh/model"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//以下为User业务方法

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	rePassword := c.PostForm("rePassword")

	user := model.User{
		Model:    gorm.Model{},
		Username: username,
		Password: password,
	}

	if rePassword != password {
		fmt.Println("两次密码不一致！")
		c.HTML(http.StatusOK, "register.html", "两次密码输入不一致！")
	} else {
		fmt.Println("注册成功！")
		dao.Mgr.Register(&user)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func GoRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", "go register")
}

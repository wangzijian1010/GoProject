package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("D:\\gopro\\GoWeb\\GoWeb1\\templates\\*")
	r.Static("/static", "D:\\gopro\\GoWeb\\GoWeb1\\static")

	// 连接到MySQL数据库
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 创建表（如果不存在）
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(255),
        password VARCHAR(255)
    )`)
	if err != nil {
		panic(err.Error())
	}

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	r.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 将数据插入数据库
		_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
		if err != nil {
			panic(err.Error())
		}

		c.HTML(http.StatusOK, "registersuccess.html", gin.H{
			"username": username,
			"password": password,
		})
	})

	r.Run(":8080")
}

package router

import (
	"github.com/gin-gonic/gin"
	"wzjwh/controller"
)

func Start() {
	r := gin.Default()

	//加载静态资源文件
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", controller.Index)

	//router group 以下是路由组
	//user路由组
	user := r.Group("user")
	{
		user.POST("/register", controller.Register)
		user.GET("/register", controller.GoRegister)
	}

	err := r.Run(":8080")
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}

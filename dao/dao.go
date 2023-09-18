package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"wzjwh/model"
)

// Manager 定义接口，接口里面用来声明操作数据库的函数
type Manager interface {
	// Register 定义用户操作的函数
	Register(user *model.User)
}

//封装数据库的db，将db变成manager，对数据库操作的db换成了manager
type manager struct {
	db *gorm.DB
}

// Mgr 声明实例方便操作
var Mgr Manager

// 初始化数据库
func init() {
	//数据库mysql用户名:mysql密码@tcp(本地端口号)/数据库名称?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:123456@tcp(127.0.0.1)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	log.Println("数据库连接成功！")
	//实例化manager，将Mgr关联到数据库，可以让Mgr对数据库进行操作
	Mgr = &manager{db: db}
	//AutoMigrate作用：让数据库自动创建表
	err = db.AutoMigrate(&model.User{}) //创建用户表
	if err != nil {
		return
	}
}

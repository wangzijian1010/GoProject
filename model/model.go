package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username,omitempty"` //首字母大写可以将类里面的字段导出供外部访问
	Password string `json:"password,omitempty"`
}

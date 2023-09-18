package dao

import "wzjwh/model"

// 以下函数为实dao中Manager接口的方法

// Register 实现用户注册
func (m *manager) Register(user *model.User) {
	m.db.Create(user)
}

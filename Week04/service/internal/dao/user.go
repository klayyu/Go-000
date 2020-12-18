package dao

import (
	"database/sql"
	"github.com/google/wire"
	"model"
	"log"
)

type UserRepository interface {
	AddUser()
}

// UserRepository接口实现
type userRepo struct {
	db *sql.DB
}

// 新增user
func (u *userRepo) AddUser() {
	user := &model.User{}
	user.Id = 1
	user.Name = "name"
	log.Println("add user :" + user.Name)
}

// 根据*sql.DB初始化 *userRepo
func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db}
}

// 声明NewUserRepo的返回值是UserRepository接口类型
var UserSet = wire.NewSet(NewUserRepo, wire.Bind(new(UserRepository), new(*userRepo)))
package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 用户表
type User struct {
	Id       int    //用户id
	Username string // 用户名
	Password string //密码
	Realname string // 真实姓名
	Avtar    int    // 头像id
}

func init() {
	orm.RegisterModel(new(User))
}

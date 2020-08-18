package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 回复表
type Reply struct {
	Id      int       // 回复id
	Content string    // 回复内容
	Time    time.Time //回复问题时间
	User    *User     `orm:"rel(fk)"` //对应回复人的id
	Ask     *Ask      `orm:"rel(fk)"` // 对应问题id
}

func init() {
	orm.RegisterModel(new(Reply))
}

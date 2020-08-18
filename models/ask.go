package models

import (
	_ "database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//提问表
type Ask struct {
	Id      int       // 问题id
	Title   string    //问题标题
	Content string    //问题内容
	Time    time.Time // 问题时间
	User    *User     `orm:"rel(fk)"` //发表问题人的id
}

func init() {
	orm.RegisterModel(new(Ask))
}

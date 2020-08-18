package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "project/routers"
)

func main() {
	beego.Run()
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/issue?loc=Local", 10)
}

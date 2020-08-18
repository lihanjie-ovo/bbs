package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"project/models"
	"strconv"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Register() {
	c.TplName = "register.html"
}

func (c *RegisterController) RegisterAct() {
	// 获取post上来的数据
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	realname := c.Input().Get("realname")
	avtar, _ := strconv.Atoi(c.Input().Get("avtar"))

	user := models.User{}
	//先验证当前用户名是否被注册
	o := orm.NewOrm()

	orm.Debug = true

	err := o.QueryTable("user").Filter("Username__eq", username).One(&user)
	//	如果err为nil表示可以查到 已存在
	if err != nil {
		newUser := models.User{}
		newUser.Username = username
		newUser.Password = password
		newUser.Realname = realname
		newUser.Avtar = avtar
		_, err := o.Insert(&newUser)

		if err != nil {
			//查到了
			c.Data["title"] = "注册失败"
			c.Data["message"] = "插入数据有误"
			c.TplName = "regmessage2.html"
			os.Exit(1)
		} else {
			c.Data["title"] = "注册成功"
			c.Data["message"] = "成功！！！"
			c.TplName = "regmessage.html"
		}
	} else {
		// 查看
		c.Data["title"] = "注册失败"
		c.Data["message"] = "已有用户注册"
		c.TplName = "regmessage2.html"
	}
}

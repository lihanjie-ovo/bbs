package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"project/models"
	"time"
)

type AddControllers struct {
	beego.Controller
}

func (c *AddControllers) List() {
	// 声明一个容器来放留言
	var askData []models.Ask
	orm.Debug = true

	// 声明一个orm模型
	o := orm.NewOrm()
	// 获取所有的数据
	o.QueryTable("ask").RelatedSel().All(&askData)

	//把数据放到模板中
	c.Data["msg"] = askData
	c.TplName = "add.html"
}

func (c *AddControllers) Add() {
	// 获取post
	title := c.Input().Get("title")
	content := c.Input().Get("content")

	// 创建留言
	msg := models.Ask{}
	msg.Title = title
	msg.Content = content
	msg.Time = time.Now()
	// 关联user
	msg.User = &models.User{Id: int(c.GetSession("userId").(int))}

	//创建orm模型
	o := orm.NewOrm()
	_, err := o.Insert(&msg)
	if err != nil {
		c.Data["title"] = "留言失败"
		c.Data["message"] = "留言插入失败"
		c.TplName = "message.html"
		os.Exit(1)
	} else {
		c.Redirect("/discuss", 302)
	}
}


package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"project/models"

)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index() {
	// 获取登录传过来的值
	name := c.GetSession("userName")
	id := c.GetSession("userId")
	av := c.GetSession("av")
	c.Data["name"] = name
	c.Data["id"] = id
	c.Data["av"] = av

	//从数据库里拿取数据
	// 声明orm模型
	o := orm.NewOrm()
	// 装载你提出的问题的切片
	var issueData []models.Ask
	//把所有数据都放入到切片中
	_, err := o.QueryTable("ask").RelatedSel().Filter("User", id).All(&issueData)
	if err != nil {
		log.Panic("有误")
	}
	c.Data["iss"] = issueData
	c.TplName = "register.html"

	// 回答表
	a := orm.NewOrm()
	var repData [] models.Reply
	_, e := a.QueryTable("reply").RelatedSel().Filter("User", id).All(&repData)
	if e != nil {
		log.Println("有误")
	}
	c.Data["rep"] = repData
	c.TplName = "register.html"
}

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"project/models"
)
type DiscussControllers struct {
	beego.Controller
}

func (c *DiscussControllers) Discuss() {
	// 声明orm模型
	o := orm.NewOrm()
	//声明一个装载所有论坛的东西
	var dcsData []models.Ask
	//把所有数据都放到切片中
	_, err := o.QueryTable("ask").RelatedSel().All(&dcsData)
	if err != nil {
		log.Panic("有误")
	}
	c.Data["dcs"] = dcsData
	c.TplName = "discuss.html"
}

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"project/models"
)

type ExamineControllers struct {
	beego.Controller
}

// 这个查看回复
func (c *ExamineControllers) Examine() {
	wt_id, _ := c.GetInt("id")
	c.SetSession("wt_id", wt_id)
	qid := c.GetSession("wt_id").(int) // 问题id
	//从数据库里拿取数据
	// 声明orm模型
	o := orm.NewOrm()
	// 装载你提出的问题的切片
	var exData []models.Reply
	//把所有数据都放入到切片中
	_, err := o.QueryTable("reply").RelatedSel().Filter("Ask", qid).All(&exData)
	if err != nil {
		log.Panic("有误")
	}
	c.Data["ex"] = exData
	c.TplName = "examine.html"
}



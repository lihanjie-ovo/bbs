package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"project/models"
	"time"
)

type ReplyControllers struct {
	beego.Controller
}

// 回答
func (c *ReplyControllers) Reply() {
	wt_id, _ := c.GetInt("id")
	c.SetSession("wt_id", wt_id)
	c.TplName = "reply.html"
}

// 回答入库的操作
func (c *ReplyControllers) ReplyAdd() {
	// 获取post
	content := c.Input().Get("content")
	uid := c.GetSession("userId").(int) // 用户id
	qid := c.GetSession("wt_id").(int)  // 问题id
	num := models.Reply{
		Content: content,
		User: &models.User{
			Id: uid,
		},
		Ask: &models.Ask{
			Id: qid,
		},
		Time: time.Now(),
	}
	_, err := orm.NewOrm().Insert(&num)
	if err != nil {
		fmt.Print(err)
	}
	c.Redirect("/discuss", 302)
}

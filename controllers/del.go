package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"project/models"
)

type DelControllers struct {
	beego.Controller
}

// 删除
func (c *DelControllers) Del() {
	uid := c.GetSession("userId").(int) // 登录post过来的Userid
	fmt.Println(uid)


	wt_id, _ := c.GetInt("id")
	c.SetSession("wt_id", wt_id)
	qid := c.GetSession("wt_id").(int) // 问题id
	d := models.Ask{
		Id: qid,
	}
	orm.NewOrm().Delete(&d, "Id")
	c.Redirect("/discuss", 302)

}

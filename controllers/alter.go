package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"project/models"
)

type AlterControllers struct {
	beego.Controller
}

// 修改
func (c *AlterControllers) Alter() {
	c.TplName = "alter.html"
}

func (c *AlterControllers) AlterAA() {
	passworda := c.GetSession("password") // 获取登录的密码
	id := c.GetSession("userId").(int)        // 获取登录时候的id

	passwordb := c.Input().Get("password")    // 获取修改密码输入时的密码
	newpass := c.Input().Get("newpass")       // 获取新的密码
	affirmpass := c.Input().Get("affirmpass") // 获取确认密码

	if passworda == passwordb && newpass == affirmpass && newpass != passworda {
		num := models.User{
			Id:       id,
			Password: newpass,
		}
		orm.NewOrm().Update(&num, "Password")
		c.Redirect("/user/login", 302)
	}


}

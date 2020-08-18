package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/http"
	"project/models"
)

type LoginController struct {
	beego.Controller
}

//显示登录页面

func (c *LoginController) Login() {
	c.TplName = "login.html"
}

//查询登录用户和密码的方法
func (c *LoginController) LoginAct() {
	//获取post的数据
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	//生成orm模型
	o := orm.NewOrm()

	orm.Debug = true
	//绑定一个用户
	user := models.User{}
	err := o.QueryTable("user").Filter("Username__eq", username).One(&user)
	if err != nil {
		c.Data["title"] = "登录失败"
		c.Data["message"] = "当前用户不存在，请重新登录"
		c.TplName = "regmessage.html"
	} else {
		//验证当前用户是否与密码一致
		if password != user.Password {
			c.Data["title"] = "登录失败"
			c.Data["message"] = "密码不正确，请重新登录"
			c.TplName = "regmessage.html"
		} else {
			//写入session记录登录状态
			c.SessionRegenerateID()
			c.SetSession("loginStatus", "ok")
			c.SetSession("userId", user.Id)
			c.SetSession("userName", user.Username)
			c.SetSession("av", user.Avtar)
			c.SetSession("password", user.Password)
			c.Redirect("/index", http.StatusFound)
		}
	}
}
func (c *LoginController) Info() {
	//获取session
	username := c.GetSession("userName").(string)
	userid := fmt.Sprintf("%d", c.GetSession("userId").(int))

	userInfo := "用户名:" + username + "，用户id=" + userid

	c.Data["title"] = "用户信息"
	c.Data["message"] = userInfo
	c.TplName = "message.html"
}

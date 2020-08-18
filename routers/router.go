package routers

import (
	"github.com/astaxie/beego"
	"project/controllers"

)

func init() {
	// 注册页面
	beego.Router("/", &controllers.RegisterController{}, "get:Register;post:RegisterAct")

	// 登录
	beego.Router("/user/login", &controllers.LoginController{}, "get:Login;post:LoginAct")
	beego.Router("/user/info", &controllers.LoginController{}, "get:Info")

	//个人主页
	beego.Router("/index", &controllers.IndexController{}, "get:Index")

	// 修改密码
	beego.Router("/alter", &controllers.AlterControllers{}, "get:Alter;post:AlterAA")

	//添加留言
	beego.Router("/add", &controllers.AddControllers{}, "get:List;post:Add")

	// 添加回复
	beego.Router("/reply", &controllers.ReplyControllers{}, "get:Reply;post:ReplyAdd")

	// 查看回复
	beego.Router("/examine", &controllers.ExamineControllers{}, "get:Examine")

	// 论坛页面
	beego.Router("/discuss", &controllers.DiscussControllers{}, "get:Discuss")

	//删除问题
	beego.Router("/del", &controllers.DelControllers{}, "get:Del")
}

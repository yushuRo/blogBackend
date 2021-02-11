package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// 用户路由
	beego.Router("/member/login", &controllers.MemberController{}, "Post:Login")
	beego.Router("/member/register", &controllers.MemberController{}, "Post:Register")

	// 首页路由
	beego.Router("/index/index", &controllers.IndexController{}, "Get:Index")

}

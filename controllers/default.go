package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	user, err := models.IdGetUser(1)
	if err != nil {
		panic(err)
	}
	//sprintf := fmt.Sprintf("%v", user)
	//for _, v := range sprintf{
	c.Ctx.WriteString(string(user.Id))
	//}
	//c.GetControllerAndAction()
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
}

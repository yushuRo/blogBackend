package controllers

type IndexController struct {
	AuthController
}

func (i *IndexController) Index() {
	i.Ctx.WriteString("主页")
}

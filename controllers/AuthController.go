package controllers

import (
	"blog/utils"
	"encoding/json"
	"strings"
)

type AuthController struct {
	BaseController
}

func (i *IndexController) Prepare() {
	token := utils.Base64Decode(i.GetString("token"))
	maps := make(map[string]string)
	_ = json.Unmarshal([]byte(strings.Replace(token, "\\", "", -1)), &maps)
	res := i.BaseController.getUserInfo2Session(maps["username"], maps["token"])
	if res != true {
		i.jsonResult(401, "无效的登录!", "")
	}
}

package controllers

import (
	"blog/models"
	"blog/utils"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// 验证用户
func (b *BaseController) getUserInfo2Session(username, token string) bool {
	user, _ := models.UsernameGetUser(username)
	if user == nil {
		return false
	}
	sessionKey := utils.Str2md5(username + user.Salt)
	checkToken := b.GetSession(sessionKey)
	if checkToken != token {
		return false
	}
	return true
}

// 设置token
func (b *BaseController) SetToken(username string) interface{} {
	user, _ := models.UsernameGetUser(username)
	if user == nil {
		return false
	}
	sessionKey := utils.Str2md5(username + user.Salt)
	token := utils.GetRandomString(50)
	b.SetSession(sessionKey, token)
	return token
}

// 过期token
func (b *BaseController) UnsetToken(username string) bool {
	user, _ := models.UsernameGetUser(username)
	if user == nil {
		return false
	}
	sessionKey := utils.Str2md5(username + user.Salt)
	b.DelSession(sessionKey)
	return true
}

// json
func (b *BaseController) jsonResult(code int64, msg string, obj interface{}) {
	r := make(map[string]interface{})
	r["code"] = code
	r["msg"] = msg
	r["obj"] = obj
	b.Data["json"] = r
	b.ServeJSON()
	b.StopRun()
}

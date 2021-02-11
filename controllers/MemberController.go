package controllers

import (
	"blog/models"
	"blog/utils"
	"encoding/json"
	"strings"
)

type MemberController struct {
	BaseController
}

// 登录
func (m *MemberController) Login() {
	username := strings.TrimSpace(m.GetString("username"))
	password := strings.TrimSpace(m.GetString("password"))
	res := models.CheckUser(username, password)
	if res == false {
		m.jsonResult(400, "用户帐号或密码不正确，请重新输入", "")
	}
	user, _ := models.UsernameGetUser(username)
	// 最新登录时间
	models.InsertLoginTime(int(user.Id))
	m.StartSession()
	token := m.SetToken(username)

	if token != false {
		enc := make(map[string]interface{})
		enc["username"] = user.Username
		enc["token"] = token
		jsonp, _ := json.Marshal(enc)
		m.jsonResult(200, "登录成功", utils.Base64Encode(string(jsonp)))
	}

}

// 注册
func (m *MemberController) Register() {
	username := strings.TrimSpace(m.GetString("username"))
	password := strings.TrimSpace(m.GetString("password"))
	member, _ := models.UsernameGetUser(username)
	if member != nil {
		m.jsonResult(400, "该帐号已注册！", "")
	}
	if password == "" {
		m.jsonResult(400, "密码不能为空！", "")
	}
	//生成盐
	salt := utils.GetRandomString(8)
	password = utils.Str2md5(password + salt)
	res := models.CreateMember(username, password, salt)
	if res == true {
		m.jsonResult(200, "注册成功！", "")
	}
	m.jsonResult(400, "网络超时，请重新注册！", "")
}

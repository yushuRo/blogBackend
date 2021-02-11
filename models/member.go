package models

import (
	"blog/utils"
	"github.com/astaxie/beego/orm"
	"time"
)

type BlogMember struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	CreateAt int64  `json:"create_at"`
	LoginAt  int64  `json:"login_at"`
}

// 验证用户信息
func CheckUser(username string, password string) bool {
	if username == "" || password == "" {
		return false
	}
	user := BlogMember{}
	err := orm.NewOrm().QueryTable("BlogMember").Filter("username", username).One(&user)
	if err != nil {
		return false
	}
	if user.Username == "" {
		return false
	}
	// 验证密码
	checkPassword := utils.Str2md5(password + user.Salt)
	if checkPassword != user.Password {
		return false
	}

	return true
}

// 通过id查找用户
func IdGetUser(id int64) (*BlogMember, error) {
	user := BlogMember{}
	err := orm.NewOrm().QueryTable("BlogMember").Filter("id", id).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 通过用户名查找用户
func UsernameGetUser(username string) (*BlogMember, error) {
	user := BlogMember{}
	err := orm.NewOrm().QueryTable("BlogMember").Filter("username", username).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 注册用户
func CreateMember(username string, password string, salt string) bool {
	var user BlogMember
	user.Username = username
	user.Password = password
	user.Salt = salt
	user.CreateAt = time.Now().Unix()
	res, err := orm.NewOrm().Insert(&user)
	if err != nil || res == 0 {
		return false
	}
	return true

}

// 插入登录时间
func InsertLoginTime(id int) bool {
	var user BlogMember
	user.Id = int64(id)
	user.LoginAt = time.Now().Unix()
	res, err := orm.NewOrm().Update(&user, "LoginAt")
	if err != nil || res == 0 {
		return false
	}
	return true
}

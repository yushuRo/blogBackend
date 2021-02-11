package sysinit

import (
	"github.com/astaxie/beego"
)

// 初始化session
func initSession() {
	// 是否开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	// session过期时间
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
	// session引擎
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	// session引擎地址
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./runtime/session"
}

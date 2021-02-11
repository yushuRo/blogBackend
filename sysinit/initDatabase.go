package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//初始化数据连接
func initDatabases() {
	//  数据库类型
	dbType := beego.AppConfig.String("db_type")
	// 连接名称
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	// 数据库名称
	dbName := beego.AppConfig.String(dbType + "::db_name")
	// 连接用户名
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	// 连接密码
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	// 连接地址
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	// 连接端口
	dbPort := beego.AppConfig.String(dbType + "::db_port")
	if dbType == "mysql" {
		dbCharset := beego.AppConfig.String(dbType + "::db_charset")
		_ = orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+
			dbPort+")/"+dbName+"?charset="+dbCharset, 30)
		orm.Debug = true
		// 最大闲置连接量
		orm.SetMaxIdleConns(beego.AppConfig.String("mysql::db_alias"),
			beego.AppConfig.DefaultInt("mysql::db_max_idle", 10))
		// 最大连接量
		orm.SetMaxOpenConns(beego.AppConfig.String("mysql::db_alias"),
			beego.AppConfig.DefaultInt("mysql::db_max_active", 128))
	}

}

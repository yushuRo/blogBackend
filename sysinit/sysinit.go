package sysinit

func init() {
	// 初始化数据库
	initDatabases()
	// 初始化session
	initSession()

}

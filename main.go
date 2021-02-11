package main

import (
	_ "blog/routers"
	_ "blog/sysinit"
	"github.com/astaxie/beego"
	_ "github.com/gomodule/redigo/redis"
)

func main() {
	beego.Run()
}

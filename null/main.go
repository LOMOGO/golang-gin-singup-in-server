package main

import (
	"server/extend/conf"
	"server/extend/validata"
	"server/model"
	"server/router"
)

func main() {
	//因为后面的包需要用到数据，所以配置包需要先启动
	conf.Setup()
	model.Setup()
	validata.Setup()

	//gin启动最好放在最后
	r := router.InitRouter()
	r.Run(":8080")
}

package main

import (
	"server/extend/validata"
	"server/model"
	"server/router"
)

func main() {
	model.Setup()
	validata.Setup()

	//gin启动最好放在最后
	r := router.InitRouter()
	r.Run(":8080")
}

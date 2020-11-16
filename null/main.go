package main

import (
	"server/model"
	"server/router"
)

func main() {
	model.Setup()

	r := router.InitRouter()
	r.Run(":8080")
}

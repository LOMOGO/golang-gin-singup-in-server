package main

import (
	"server/model"
	"server/router"
)

func main() {
	router.Setup()
	model.Setup()
}

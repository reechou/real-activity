package main

import (
	"github.com/reechou/real-activity/config"
	"github.com/reechou/real-activity/controller"
)

func main() {
	controller.NewActLogic(config.NewConfig()).Run()
}

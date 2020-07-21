package app

import "github.com/arshabbir/docker1/src/controller"

func StartApplication() {

	controller := controller.NewItemsController()

	controller.Start()
}

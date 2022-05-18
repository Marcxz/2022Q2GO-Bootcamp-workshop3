package main

import (
	"fmt"

	api "architecture/api"
	controller "architecture/controller"
	database "architecture/db"
)

func main() {

	db := database.NewSQL()
	geo := api.NewGeo()
	c := controller.NewController(geo, db)

	info, _ := c.GetInfoUserById("2")
	_ = info

	fmt.Println(info)
}

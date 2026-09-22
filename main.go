package main

import (
	"log"
	_ "rental-property-api/routers"
	"rental-property-api/services"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	err := services.LoadProperties()
	if err!= nil{
		log.Fatal(
			"Failed to load properties:",
			err,
		)
	}
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

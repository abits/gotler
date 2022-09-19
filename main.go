package main

import (
	"github.com/abits/gotler/controllers"
	"github.com/abits/gotler/mappings"
)

func main() {
	controllers.InitDB()
	mappings.CreateUrlMappings()
	mappings.Router.Run(":8080")
}

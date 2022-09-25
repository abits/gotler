package main

import (
	"github.com/abits/gotler/mappings"
)

func main() {
	mappings.CreateUrlMappings()
	mappings.Router.Run(":8080")
}

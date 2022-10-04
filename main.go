package main

import (
	"github.com/abits/gotler/mappings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:gotler@tcp(127.0.0.1:3306)/gotler?charset=utf8mb4&parseTime=True&loc=Local"
	gorm.Open(mysql.Open(dsn), &gorm.Config{})
	mappings.CreateUrlMappings()
	mappings.Router.Run(":8080")
}

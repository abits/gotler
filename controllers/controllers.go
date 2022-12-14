package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func GetUserDetail(c *gin.Context) {
	c.String(200, "Success")
}

func GetUser(c *gin.Context) {
	c.String(200, "Success")
}

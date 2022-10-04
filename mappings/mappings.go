package mappings

import (
	"github.com/abits/gotler/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	Router.Use(controllers.Cors())

	Router.GET("/users/:id", controllers.GetUserDetail)
	Router.GET("/users", controllers.GetUser)

}

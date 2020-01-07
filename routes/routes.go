package routes

import (
	"auth/controller"

	"github.com/gin-gonic/gin"
)

// Initialize ....
func Initialize() {
	r := gin.Default()

	r.GET("/test", controller.Home)
	r.POST("/create", controller.Createuser)
	r.POST("/article", controller.Createarticle)
	r.GET("/users", controller.Getuserbyid)
	r.GET("/getArticles", controller.Getarticle)
	r.DELETE("/delete", controller.Deleteuser)
	r.PUT("/update", controller.Updateuser)

	r.Run()
}

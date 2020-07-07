package routes

import (
	"gin/controller"

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
	r.POST("/todo", controller.HandleCreateTodo)
	r.GET("/GetTodo", controller.HandleGetAllTodo)
	r.GET("/GetTodo/:id", controller.HandleGetTodoById)
	//r.PAtCH("/updateTodo", controller.HandleUpdateTodo)
	r.POST("/login", controller.Login)


	r.Run()
}

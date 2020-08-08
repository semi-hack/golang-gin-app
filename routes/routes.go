package routes

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

// Initialize ....
func Initialize() {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.POST("/create", controller.Createuser)
		user.GET("/users", controller.Getuserbyid)
		user.PATCH("/update", controller.Updateuser)
		user.DELETE("/delete", controller.Deleteuser)
	}
	
	r.GET("/test", controller.Home)
	r.POST("/article", controller.Createarticle)
	r.GET("/getArticles", controller.Getarticle)
	r.POST("/todo", controller.HandleCreateTodo)
	r.GET("/GetTodo", controller.HandleGetAllTodo)
	r.GET("/GetTodo/id", controller.HandleGetTodoByID)
	//r.PAtCH("/updateTodo", controller.HandleUpdateTodo)
	r.POST("/login", controller.Login)


	r.Run()
}

package controller

import (
	"gin/services"
	"gin/models"
	"log"
	"github.com/gin-gonic/gin"
)

// HandleCreateTodo ...
func HandleCreateTodo(c *gin.Context)  {
	var todo models.Todo


	if err:= c.ShouldBindJSON(&todo); err != nil {
		log.Print(err)
		c.JSON(400, gin.H{"msg": err })
		return
	}
	id, err := services.Create(&todo)
	if err != nil {
		c.JSON(500, gin.H{"msg": err })
	}
	c.JSON(200, gin.H{"id": id })
	
}

// HandleGetAllTodo ...
func HandleGetAllTodo(c *gin.Context) {

	T, err := services.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
	}
	c.JSON(200, gin.H{"Todos": T})
}

// HandleGetTodoByID ...
func HandleGetTodoByID(c *gin.Context) {
	//var todo models.Todo
	
	id := struct {
		ID string `form:"_id"`
	}{}
	
	if err:= c.ShouldBindQuery(&id); err != nil {
		c.JSON(400, gin.H{"msg": err })
		return
	}
	var T, err = services.GetTodoByID(id.ID)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
	}
	c.JSON(200, gin.H{"Todo": T})
}

// HandleUpdateTodo ...
// func HandleUpdateTodo(c *gin.Context) {
// 	var todo models.Todo
// 	id:= c.Param("id")
// 	if err := c.ShouldBindJSON(&id); err != nil {
// 		log.Print(err)
// 		c.JSON(400, gin.H{"msg": err})
// 		return
// 	}
// 	savedTodo, err := services.Update(&todo)
// 	if err != nil {
// 		c.JSON(500, gin.H{"msg": err})
// 		return
// 	}
// 	c.JSON(200, gin.H{"Todo": savedTodo})
// }
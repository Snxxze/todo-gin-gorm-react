package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	todo := r.Group("/todos")
	todo.Use(middlewares.AuthMiddleware())
	{
		todo.GET("/", controllers.GetTodos)
		todo.POST("/", controllers.CreateTodo)
		todo.PUT("/:id", controllers.UpdateTodo)
		todo.DELETE("/:id", controllers.DeleteTodo)
	}
}
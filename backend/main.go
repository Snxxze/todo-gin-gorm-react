package main

import (
	"backend/configs"
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main(){
	configs.ConnectionDB()
	configs.SetupDatabase()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
  	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    AllowCredentials: true,
  }))
	
	routes.RegisterRoutes(r)

	if err := r.Run(":8000"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
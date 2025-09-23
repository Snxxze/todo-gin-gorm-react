package main

import (
	"backend/configs"
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	configs.ConnectionDB()
	configs.SetupDatabase()

	r := gin.Default()

	routes.RegisterRoutes(r)

	if err := r.Run(":8000"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hng-stage3-task-automated-email-service/handlers"
	"hng-stage3-task-automated-email-service/middleware"
	"log"
)

func main() {

	r := gin.Default()

	r.Use(middleware.SetUpCORS())

	r.POST("/target_url", handlers.LoginNoOauthHandler)
	r.GET("/integration.json", handlers.ReturnIntegrationJSON)

	fmt.Println("Server running on port 8080...")
	log.Fatal(r.Run(":8080"))
}

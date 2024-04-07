package main

import (
	"fmt"
	"goapi/controllers"
	"goapi/initializers"
	"goapi/middleware"

	"github.com/gin-gonic/gin"
	//"initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	fmt.Println("Hello")
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}

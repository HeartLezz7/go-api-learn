package main

import (
	"fmt"
	"go-api/src/model"
	"go-api/src/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	engine := gin.Default()

	port := os.Getenv("PORT")

	r := model.Server{Engine: engine}
	router.UserRouter(r, "/user")
	router.TodoRouter(r, "/todo")
	fmt.Println("server run on port ", port)
	r.Engine.Run(":" + port)
}

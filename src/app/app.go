package main

import (
	"fmt"
	"go-api/src/model"
	"go-api/src/router"

	// "database/sql"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	port := ":8080"
	r := model.Server{Engine: engine}
	router.UserRouter(r, "/user")
	fmt.Println("server run on port ", port)
	r.Engine.Run(port)
}

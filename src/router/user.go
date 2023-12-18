package router

import (
	"go-api/src/controllers"
	"go-api/src/model"
)

func UserRouter(server model.Server, PATH string) {
	server.Engine.GET(PATH, controllers.GetUser)
}

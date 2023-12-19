package router

import (
	"go-api/src/controllers"
	"go-api/src/model"
)

func UserRouter(server model.Server, PATH string) {
	userRouter := server.Engine.Group(PATH)

	userRouter.GET("/getAll", controllers.GetAllUserId)
	userRouter.POST("create", controllers.CreateUser)
}

func TodoRouter(server model.Server, PATH string) {
	todoRouter := server.Engine.Group(PATH)

	todoRouter.GET("getList")
}

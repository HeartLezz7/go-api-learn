package router

import (
	"go-api/src/controllers"
	"go-api/src/model"
	"go-api/src/repo"
)

func UserRouter(server model.Server, PATH string) {
	userRouter := server.Engine.Group(PATH)

	db := repo.Database()

	userController := controllers.NewUserController(db)

	userRouter.GET("/getAll", userController.GetAllUserId)
	userRouter.POST("create", userController.CreateUser)
	userRouter.PATCH("/update/:id", userController.UpdatePassword)
	userRouter.DELETE("/delete/:id", userController.DeleteUser)
}

func TodoRouter(server model.Server, PATH string) {
	todoRouter := server.Engine.Group(PATH)

	todoRouter.GET("getList")
}

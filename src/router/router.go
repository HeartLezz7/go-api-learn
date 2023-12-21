package router

import (
	"go-api/src/controllers/todo_controller"
	"go-api/src/controllers/user_controller"
	"go-api/src/model"
	"go-api/src/repo"
)

func UserRouter(server model.Server, PATH string) {

	userRouter := server.Engine.Group(PATH)

	db := repo.Database()

	userController := user_controller.NewUserController(db)

	userRouter.GET("/getAll", userController.GetAllUserId)
	userRouter.POST("create", userController.CreateUser)
	userRouter.PATCH("/update/:id", userController.UpdatePassword)
	userRouter.DELETE("/delete/:id", userController.DeleteUser)
}

func TodoRouter(server model.Server, PATH string) {

	todoRouter := server.Engine.Group(PATH)

	db := repo.Database()

	todoController := todo_controller.NewTodoController(db)

	todoRouter.GET("getList", todoController.GetAllTodos)
}

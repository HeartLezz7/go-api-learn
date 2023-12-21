package todo_controller

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	GetAllTodos(*gin.Context)
}
type todoController struct {
	db *sql.DB
}

func NewTodoController(db *sql.DB) TodoController {
	return todoController{db: db}
}

func (t todoController) GetAllTodos(c *gin.Context) {

	data, error := t.db.Query("SELECT id,title,status,userId FROM todoList")

	if error != nil {
		c.JSON(500, "Invalid")
		return
	}
	defer data.Close()

	type todoLists struct {
		Id     int
		Titles string
		Status bool
		UserId int
	}

	var todos todoLists

	todo := []todoLists{}

	for data.Next() {
		error := data.Scan(&todos.Id, &todos.Titles, &todos.Status, &todos.UserId)
		if error != nil {
			fmt.Println(error.Error())
		}
		todo = append(todo, todos)

	}

	c.JSON(200, todo)
}

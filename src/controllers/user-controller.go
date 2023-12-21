package controllers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUserId(*gin.Context)
	CreateUser(*gin.Context)
	UpdatePassword(*gin.Context)
	DeleteUser(*gin.Context)
}

type userController struct {
	db *sql.DB
}

func NewUserController(db *sql.DB) UserController {
	return userController{db: db}
}

func (u userController) GetAllUserId(res *gin.Context) {
	// db := repo.Database()

	getAll := "select id from user ORDER BY id"
	result, err := u.db.Query(getAll)
	if err != nil {
		fmt.Println("ERROR", err.Error())
		return
	}
	defer result.Close()
	ids := []int{}
	for result.Next() {
		var id int
		err := result.Scan(&id) // ส่ง แอดเดรสไปเช็ค
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}
	res.JSON(200, ids)
}

func (u userController) CreateUser(c *gin.Context) {
	// db := repo.Database()
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	errorRequest := c.BindJSON(&data)
	fmt.Printf("data %+v\n", data)

	if errorRequest != nil {
		fmt.Printf("this is error %s\n", errorRequest)
	}

	create, _ := u.db.Prepare("INSERT INTO user (username, password, email) VALUES (?, ?, ?)")
	result, createError := create.Exec(data.Username, data.Password, data.Email)
	fmt.Println("This is result", result)

	if createError != nil {
		fmt.Println(createError)
		c.JSON(500, createError)
	}
	defer create.Close()

	num, _ := result.RowsAffected()
	numLast, _ := result.LastInsertId()
	fmt.Println(num, numLast)

	c.JSON(201, data)
}

func (u userController) UpdatePassword(c *gin.Context) {

	// db := repo.Database()

	data := make(map[string]string)
	error := c.BindJSON(&data)
	fmt.Println(error)

	params := c.Params
	id := params[0].Value

	key := make([]string, 0, 2)
	value := make([]interface{}, 0, 2)
	for k, v := range data {
		key = append(key, k)
		value = append(value, v)
	}
	value = append(value, id)

	keyString := strings.Join(key, " = ?, ")

	rawString := fmt.Sprintf("UPDATE user SET %s = ? WHERE id = ? ", keyString)

	prepare, _ := u.db.Prepare(rawString)
	defer prepare.Close()

	result, updateError := prepare.Exec(value...)
	fmt.Println(result)
	fmt.Println("CHECK ERROR", updateError)
	if updateError != nil {
		c.JSON(500, updateError)
	}

	c.JSON(200, result)
}

func (u userController) DeleteUser(c *gin.Context) {
	// db := repo.Database()
	params := c.Params
	id := params[0].Value

	prepare, _ := u.db.Prepare("DELETE FROM user where id = ? ")
	result, deleteError := prepare.Exec(id)
	if deleteError != nil {
		fmt.Println(deleteError.Error())
		c.JSON(500, deleteError)
	}

	c.JSON(200, result)
}

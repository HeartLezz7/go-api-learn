package controllers

import (
	"fmt"
	"go-api/src/repo"

	"github.com/gin-gonic/gin"
)

func GetAllUserId(res *gin.Context) {
	db := repo.Database()
	getAll := "select id from user ORDER BY id"
	result, err := db.Query(getAll)
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

func CreateUser(c *gin.Context) {
	// req := c.Request.Body
	db := repo.Database()
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

	create, _ := db.Prepare("INSERT INTO user (username, password, email) VALUES (?, ?, ?)")
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

func UpdateUser(c *gin.Context) {
	id := c.Params
	fmt.Printf("check params %+v\n", id[0].Value)
	c.JSON(200, id)
}

func DeleteUser(c *gin.Context) {
	db := repo.Database()
	params := c.Params
	id := params[0].Value

	prepare, _ := db.Prepare("DELETE FROM user where id = ? ")
	result, deleteError := prepare.Exec(id)
	if deleteError != nil {
		fmt.Println(deleteError.Error())
		c.JSON(500, deleteError)
	}

	c.JSON(200, result)
}

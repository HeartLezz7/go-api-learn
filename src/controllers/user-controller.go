package controllers

import (
	"fmt"
	"go-api/src/repo"

	"github.com/gin-gonic/gin"
)

func GetAllUserId(res *gin.Context) {
	db := repo.Database()
	getAll := "select id from user"
	fmt.Println("BEFORE QUERY", db)
	result, err := db.Query(getAll)
	fmt.Println(result, "CHECK REUSLT")
	if err != nil {
		fmt.Println("ERROR", err.Error())
		return
	}
	defer result.Close()
	ids := []int{}
	for result.Next() {
		var id int
		fmt.Println(id, "BEFORE")
		err := result.Scan(&id) // ส่ง แอดเดรสไปเช็ค
		fmt.Println(id, "AFTER")
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}
	res.JSON(200, ids)
}

func CreateUser(res *gin.Context) {
	res.JSON(201, nil)
}

package controllers

import (
	"fmt"
	"go-api/src/repo"

	"github.com/gin-gonic/gin"
)

func GetAllUserId(c *gin.Context) {
	db := repo.Database()
	getAll := "select id from user"
	fmt.Println("BEFORE QUERY", db)
	result, err := db.Query(getAll)
	fmt.Println(result, "CHECK REUSLT")
	if err != nil {
		fmt.Println(err, "CHECK ERROR")
		panic(err)
	}
	defer result.Close()
	fmt.Println("AFTER QUERY")
	if err != nil {
		fmt.Println("ERROR", err.Error())
		return
	}
	ids := []int{}
	for result.Next() {
		var id int
		err := result.Scan(&id)
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}
	c.JSON(200, ids)

}

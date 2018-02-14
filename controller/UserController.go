package controller

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"restfulUser/db"
)

func GetUser(c *gin.Context) {
	var sql =  new (db.Mysql)
	sql.Sql = "select * from users"

	result, err := db.Db.Query(sql)
	fmt.Println(result)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func GetUsers(c *gin.Context) {
	var sql =  new (db.Mysql)
	sql.Sql = "select * from users"

	result, err := db.Db.Query(sql)
	fmt.Println(result)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

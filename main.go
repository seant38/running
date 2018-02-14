package main

import (
	"fmt"
	"encoding/base64"
	"strings"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"

	"restfulUser/db"
	"restfulUser/controller"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func checkAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := strings.SplitN(c.GetHeader("Authorization"), " ", 2)
		if len(s) != 2 {
			respondWithError(401, "Unauthorized", c)
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])

		//sEnc := base64.StdEncoding.EncodeToString([]byte(b))

		if err != nil {
			respondWithError(401, "Unauthorized", c)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			respondWithError(401, "Unauthorized", c)
			return
		}

		//password := "secret"
		//hash, _ := HashPassword(password) // ignore error for the sake of simplicity
		//
		//fmt.Println("Password:", password)
		//fmt.Println("Hash:    ", hash)
		//
		//match := CheckPasswordHash(password, hash)
		//fmt.Println("Match:   ", match)

		if pair[0] != "user" || pair[1] != "pass" {
			respondWithError(401, "Unauthorized", c)
			return
		}

		c.Next()
	}
}

func main() {
	db.Init()
	router := gin.Default()
	router.Use(checkAuth())
	router.Use(CORSMiddleware())

	router.GET("/user/:id", controller.GetUser)
	router.GET("/user/", controller.GetUsers)

	//// POST new person details
	//router.POST("/person", func(c *gin.Context) {

	//// PUT - update a person details
	//router.PUT("/person", func(c *gin.Context) {

	//// Delete resources
	//router.DELETE("/person", func(c *gin.Context) {

	router.Run(":3000")
}

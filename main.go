package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Country string `json:"country"`
}

func users(context *gin.Context) {
	bytes, err := ioutil.ReadFile("./data/users.json")
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	if err := json.Unmarshal(bytes, &users); err != nil {
		log.Fatal(err)
	}
	context.JSON(http.StatusOK, users)
}

func main() {
	route := gin.Default()

	route.GET("/users", users)
	route.Run()
}

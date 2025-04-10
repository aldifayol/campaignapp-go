package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID string `json:"id"`
	FullName string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email string `json:"email"`
	IsAdmin bool `json:"is_admin"`
}

var users = []user{
	{ ID: "1", FullName: "Aldi Buds", PhoneNumber: "08114341004", Email: "aldifayol@gamil.com", IsAdmin: true},
	{ ID: "2", FullName: "Lopez Hanna", PhoneNumber: "08114321234", Email: "lopez@gamil.com", IsAdmin: true},
}

func getUsers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users)
}

func main() {
    router := gin.Default()
    router.GET("/users", getUsers)

    router.Run("localhost:8080")
}
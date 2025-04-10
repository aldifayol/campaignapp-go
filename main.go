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

// postUsers adds a user from JSON received in the request body.
func postUsers(c *gin.Context) {
    var newUser user

    // Call BindJSON to bind the received JSON to
    // newUser.
    if err := c.BindJSON(&newUser); err != nil {
        return
    }

    // Add the new album to the slice.
    users = append(users, newUser)
    c.IndentedJSON(http.StatusCreated, newUser)
}

func getUsers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users)
}

func main() {
    router := gin.Default()
    router.GET("/users", getUsers)
		router.POST("/users", postUsers)

    router.Run("localhost:8080")
}
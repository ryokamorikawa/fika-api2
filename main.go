package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// user represents data about a record user.
type user struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Adress string  `json:"adress"`
	Age    float64 `json:"age"`
}

// users slice to seed record user data.
var users = []user{
	{ID: "1", Name: "佐藤　太郎", Adress: "北海道", Age: 56},
	{ID: "2", Name: "佐川　春子", Adress: "東京", Age: 17},
	{ID: "3", Name: "佐々木　聡", Adress: "静岡", Age: 39},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUsers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	router.Run("0.0.0.0:" + port)
}

// getUsers responds with the list of all users as JSON.
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// postUsers adds an user from JSON received in the request body.
func postUsers(c *gin.Context) {
	var newuser user

	// Call BindJSON to bind the received JSON to
	// newuser.
	if err := c.BindJSON(&newuser); err != nil {
		return
	}

	// Add the new user to the slice.
	users = append(users, newuser)
	c.IndentedJSON(http.StatusCreated, newuser)
}

// getUserByID locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func getUserByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of users, looking for
	// an user whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/getUserDetails", getUserDetails)

	router.Run("localhost:8081")
}

func getUserDetails(c *gin.Context) {

	jsondata, _ := json.Marshal(`{
		"Id": 1,
		"Name": "satyam",
		"EmailId": "satya@xyz.co",
		"Address": "GopalNagar, India"
	}`)
	log.Println("JsonData is: ", jsondata)

	c.JSON(http.StatusOK, `{"Id": 1, "Name": "satyam", "EmailId": "satya@xyz.co", "Address": "GopalNagar, India"}`)
}

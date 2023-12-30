package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mymiddlewares/middlewares"
)

type Student struct {
	Id   int
	Name string
}

func main() {

	router := gin.Default()
	//router.Use(middlewares.Authenticate) // this will be applied at the app level.

	//router.GET("/students", middlewares.Authenticate, getStudents) // at the api level
	router.GET("/students", middlewares.Authenticate(), getStudents) // at the api level function which returns function
	router.GET("/studentswithquery", getStudentsWithQueryString)
	router.GET("/students/:id", getStudentsWithParams)
	router.POST("/students", addStudent)

	router.Run()
}

func getStudents(c *gin.Context) {

	students := []Student{
		{1, "Piyush"},
		{2, "Satish"},
	}

	c.JSON(200, gin.H{"data": students})
}

func getStudentsWithQueryString(c *gin.Context) {

	firstNameStr := c.Query("firstName")
	lastNameStr := c.Query("lastName")

	c.JSON(200, gin.H{
		"firstname": firstNameStr,
		"lastname":  lastNameStr,
	})
}

func getStudentsWithParams(c *gin.Context) {

	id := c.Param("id")

	c.JSON(200, gin.H{"id": id})
}

func addStudent(c *gin.Context) {

	body := c.Request.Body

	var student Student
	json.NewDecoder(body).Decode(&student)

	c.JSON(201, gin.H{"added": student})
}

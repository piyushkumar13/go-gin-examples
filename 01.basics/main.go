package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Student struct {
	Id   int
	Name string
}

func main() {

	//router := gin.Default() // it attaches the default logger and server config.

	router := gin.New() // this does not attaches the default logger. But we can add the logger using use keyword.

	router.Use(gin.Logger())

	router.GET("/students", getStudents)
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

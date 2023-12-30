package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"loggingwithgin/logger"
	"os"
)

type Student struct {
	Id   int
	Name string
}

func main() {

	router := gin.Default()

	loggingfile, _ := os.Create("logfile.log")
	//gin.DefaultWriter = io.MultiWriter(loggingfile) // this will set the logging to the log file
	gin.DefaultWriter = io.MultiWriter(loggingfile, os.Stdout) // this will set the logging to the log file as well as standard output

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) { // This will format the logs of the route only.
		log.Printf("endpoint formatted information is %v %v %v %v \n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	router.Use(gin.LoggerWithFormatter(logger.FormatLogs)) // for this we can also use gin.New() instead of gin.Default() as well.
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

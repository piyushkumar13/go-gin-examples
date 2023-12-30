package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"logrusexample2/config"
	"logrusexample2/logger"
)

type Student struct {
	Id   int
	Name string
}

func init() {

	config.Init()
	logger.Init()
}

func main() {

	//logrus.SetLevel(logrus.DebugLevel)
	//logrus.SetReportCaller(true)               // it will add more details to the logs like function name and file name as well.
	//logrus.SetFormatter(&logrus.TextFormatter{ // Formatting the log in text
	//	DisableTimestamp: true,
	//	FullTimestamp:    false,
	//})
	//
	////logrus.SetFormatter(&logrus.JSONFormatter{ // Formats the log in json
	////	DisableTimestamp: true,
	////})
	//
	////logrus.SetOutput(os.Stdout) // set the output to console
	//
	//file, _ := os.Create("logfile.log")
	//logrus.SetOutput(file) // sets the output to the file
	//
	//multiWriter := io.MultiWriter(os.Stdout, file)
	//logrus.SetOutput(multiWriter) // sets the output to multiple destinations like file and console

	router := gin.New()

	logrus.Println("Hello! Example of logrus")

	router.GET("/students", getStudents)
	router.GET("/studentswithquery", getStudentsWithQueryString)
	router.GET("/students/:id", getStudentsWithParams)
	router.POST("/students", addStudent)

	router.Run()
}

func getStudents(c *gin.Context) {

	logrus.WithField("MoreInfo", "Getting Students").Info("We are getting list of students")
	logrus.WithFields(logrus.Fields{
		"field1": "value1",
		"field2": "value2",
	}).Info("Some more set of key value pairs")

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

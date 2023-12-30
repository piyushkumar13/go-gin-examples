package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Student struct {
	Id   int
	Name string
}

func main() {

	router := gin.Default()

	//router.GET("/students", getStudents)
	//router.GET("/studentswithquery", getStudentsWithQueryString)
	//router.GET("/students/:id", getStudentsWithParams)
	//router.POST("/students", addStudent)

	auth := gin.BasicAuth(gin.Accounts{
		"root": "1234",
	})

	adminGrp := router.Group("/admin")
	{
		adminGrp.GET("/students", getStudents)
		adminGrp.GET("/studentswithquery", getStudentsWithQueryString)
	}

	clientGrp := router.Group("/client")
	{
		clientGrp.GET("/students/:id", getStudentsWithParams)

	}

	authGrp := router.Group("/authgrp", auth)
	{
		authGrp.POST("/students", addStudent)
	}

	server := &http.Server{
		Handler:      router,
		Addr:         ":9090",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	server.ListenAndServe()
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

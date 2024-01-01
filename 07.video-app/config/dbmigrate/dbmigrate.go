package main

import (
	"log"
	"video-app/config"
	"video-app/domain/entity"
)

func init() {

	config.LoadEnvs()
	config.ConnectToDb()
}
func main() {

	err := config.DB.AutoMigrate(&entity.Video{}, &entity.Person{})

	if err != nil {
		log.Fatal("Cannot dbmigrate db", err.Error())
	}
}

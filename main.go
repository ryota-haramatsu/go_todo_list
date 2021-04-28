package main

import (
	"fmt"
	"log"

	"example.com/go_todoapp/app/models"
)

func main() {
	/*
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		log.Println("test")
	*/
	// fmt.Println(models.Db)
	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.Password = "testtest"
	fmt.Println(u)
	log.Println(u)

	u.CreateUser()
}

package main

import (
	"fmt"
	"log"

	"example.com/go_todoapp/config"
)

func main() {
	log.Println("test")
	fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
}

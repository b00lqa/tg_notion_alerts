package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("vim-go")

	// Reding config.yml file.
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	nc := createNotionClient()
	db, err := nc.getDB()
	if err != nil {
		log.Fatalf("error getting notion db: %v", err)
	}

	taskList, err := nc.getTasksList(db)
	if err != nil {
		log.Fatalf("error getting tasks: %v", err)
	}

	for _, task := range taskList {
		log.Print(task.Properties["Name"])
	}

	log.Print(db.Title)
}

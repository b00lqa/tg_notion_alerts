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
	db, err := nc.GetDB()
	if err != nil {
		log.Fatalf("error getting notion db: %v", err)
	}

	log.Print(db.Title)
}

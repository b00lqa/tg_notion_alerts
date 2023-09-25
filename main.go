package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("vim-go")
	err := readConfig()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	log.Println(
		viper.Get("Telegram").(map[string]interface{})["api"],
	)
}

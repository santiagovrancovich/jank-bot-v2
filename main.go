package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("config.json")

	if err != nil {
		log.Fatalln("Missing or invalid config file")
	}

	var conf Config
	err = json.Unmarshal(file, &conf)

	fmt.Println(conf)
}

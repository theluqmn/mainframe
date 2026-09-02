package main

import (
	"fmt"
	"os"
	"encoding/json"
	"log"
	"github.com/labstack/echo/v4"

	"main/utils"
)

func loadConfiguration() map[string]any {
	var config map[string]any
	configPath := "./config.json"
	
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalf("error decoding config: %v", err)
	}
	log.Println("config loaded")
	return config
}

func main() {
	fmt.Print("\033[H\033[2J")
	log.Println("gello world!")
	
	config := loadConfiguration()

	// initialise database
	data := "./data"
	if configData, ok := config["data"]; ok {
		data = fmt.Sprintf("%v", configData)
	}
	
	db, err := utils.NewDB(fmt.Sprintf("%s/database.sqlite", data))
	if err != nil {
		log.Fatalf("error initialising database: %v", err)
	}
	defer db.Close()
	
	// initialise echo server
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// start server
	port := "8080"
	if configPort, ok := config["port"]; ok {
		port = fmt.Sprintf("%v", configPort)
	}

	log.Printf("starting at port :%s", port)
	err = e.Start(":"+port)
	if err != nil {
		log.Fatalf("shut down unexpectedly: %v", err)
	}
}
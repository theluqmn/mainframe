package main

import (
	"fmt"
	"os"
	"encoding/json"
	"log"
	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("gello world!")
	
	// fetch configuration
	configPath := "./config.json"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}
	log.Println("loading config:", configPath)

	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	var config map[string]any
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalf("error decoding config: %v", err)
	}
	log.Println("config loaded")
	
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
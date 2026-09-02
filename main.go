package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"io"
	"html/template"

	"github.com/labstack/echo/v4"

	"main/handlers"
	"main/utils"
)

//go:embed web
var webFiles embed.FS

type Template struct { templates *template.Template }

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

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

func registerRoutes(e *echo.Echo) {
	webHandler := handlers.NewWebHandler()

	e.GET("/", webHandler.RenderDesktop)
	e.Static("/static", "static")
	
	log.Println("all routes registered")
} 

func main() {
	fmt.Print("\033[H\033[2J")
	log.Println("gello world!")

	renderer := &Template{ templates: template.Must(template.ParseFS(os.DirFS("web"), "*.html", "*/*.html")) }
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
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.IPExtractor = echo.ExtractIPFromRealIPHeader()
	e.Renderer = renderer

	registerRoutes(e)

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
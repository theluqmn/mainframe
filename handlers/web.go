package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WebHandler struct{}

func NewWebHandler() *WebHandler {
	return &WebHandler{}
}

func (h *WebHandler) RenderDesktop(c echo.Context) error {
	log.Printf("[%s] GET /screen", c.RealIP())
	
	return c.Render(http.StatusOK, "desktop", map[string]any{})
}
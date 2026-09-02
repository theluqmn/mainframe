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

func (h *WebHandler) RenderHome(c echo.Context) error {
	log.Printf("[%s] GET /", c.RealIP())

	return c.Render(http.StatusOK, "home", map[string]any{})
}
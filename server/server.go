package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/jpw547/pyxis/server/handlers"
	"github.com/labstack/echo"
)

func main() {
	port := ":13000"

	router := echo.New()

	// endpoints

	router.GET("/food", handlers.SearchFoodPlaces)
	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	err := router.StartServer(&server)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("failed to start server: %s", err)
	}
}

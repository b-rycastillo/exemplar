package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/example/exemplar/internal/handler"
	"github.com/example/exemplar/internal/service"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	userService := service.NewUserService()
	handlers := handler.NewHandler(userService)
	router := handler.SetupRoutes(handlers)

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}

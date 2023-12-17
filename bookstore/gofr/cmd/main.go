// cmd/main/main.go
package main

import (
	"log"
	"net/http"

	"github.com/Fayez-2403/gofr-bookstore/pkg/routes"
	"github.com/Fayez-2403/gofr-bookstore/pkg/config"
	"github.com/gofr-dev/gofr"
)

func main() {
	app := gofr.New()

	config.Connect() // Initialize database connection

	r := app.Router()

	routes.RegisterBookStoreRoutes(r)

	// Start the GoFr server
	app.Start()
}

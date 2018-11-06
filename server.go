package main

import (
	"github.com/eaperezc/golosina/framework"
	"github.com/eaperezc/golosina/middleware"
	"github.com/eaperezc/golosina/routes"
)

func main() {

	app := framework.New()

	// Prepare routes
	routes.WebRoutes(app.Router)
	routes.APIRoutes(app.Router)

	app.Router.Use(middleware.LogRequests())

	app.Start()
}

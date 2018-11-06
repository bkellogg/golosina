package main

import (
	"github.com/eaperezc/golosina/controllers"
	"github.com/eaperezc/golosina/framework"
	"github.com/eaperezc/golosina/middleware"
)

func main() {

	app := framework.New()

	app.Router.Resource("/applications", &controllers.ApplicationController{})

	app.Router.Group("/api", func(r *framework.Router) {
		c := &controllers.ApplicationController{}
		r.Get("/applications", c.Index)
	})

	app.Router.Use(middleware.LogRequests())

	app.Start()
}

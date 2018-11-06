package routes

import (
	"github.com/eaperezc/golosina/controllers"
	"github.com/eaperezc/golosina/framework"
)

// APIRoutes will contain the list of all routes
// that are part of the API
func APIRoutes(r *framework.Router) {

	// We want to prefix the API routes
	r.Group("/api/v1", func(r *framework.Router) {
		c := &controllers.ApplicationController{}
		r.Get("/applications", c.Index)
	})
}

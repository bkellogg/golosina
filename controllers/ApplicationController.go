package controllers

import (
	"github.com/eaperezc/golosina/framework"
	"github.com/eaperezc/golosina/models"
	"github.com/eaperezc/golosina/views"
)

// ApplicationController definition for the controller
type ApplicationController struct{}

// ApplicationResponse has the structure we will respond from
// this controller. Message will be used to send error
// string when something bad happens
type ApplicationResponse struct {
	Success      bool                  `json:"success"`
	Application  *models.Application   `json:"application,omitempty"`
	Applications []*models.Application `json:"applications,omitempty"`
	Message      string                `json:"message,omitempty"`
}

// Index will return a list of all the resource
// Route: GET /applications
func (c *ApplicationController) Index(ctx *framework.Context) {
	var applications []*models.Application
	ctx.Database.Debug().Find(&applications)

	ctx.Response.JSON(&ApplicationResponse{
		Success:      true,
		Applications: applications,
	})
}

// Show the resource by id
// Route: GET /applications/{id}
func (c *ApplicationController) Show(ctx *framework.Context) {

	params, valid := ctx.Request.Validate(map[string]string{
		"id": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

	var application models.Application
	ctx.Database.Debug().First(&application, params["id"]) // find app with id 1

	if application.ID == 0 {
		ctx.Response.JSON(&ApplicationResponse{Success: false, Message: "Application doesn't exist"})
		return
	}
	ctx.Response.JSON(&ApplicationResponse{Success: true, Application: &application})
}

// Create will add a new resouce
// Route: POST /applications
func (c *ApplicationController) Create(ctx *framework.Context) {

	params, valid := ctx.Request.Validate(map[string]string{
		"name": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

	application := models.Application{
		Name: params["name"].(string),
	}

	ctx.Database.Debug().Create(&application)
	ctx.Response.JSON(&ApplicationResponse{Success: true, Application: &application})
}

// Update will change the info of the resource
// Route: PUT /applications/{id}
func (c *ApplicationController) Update(ctx *framework.Context) {

	params, valid := ctx.Request.Validate(map[string]string{
		"id":   "required",
		"name": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

	var application models.Application
	ctx.Database.Debug().First(&application, params["id"])

	if application.ID == 0 {
		ctx.Response.JSON(&ApplicationResponse{Success: false, Message: "Application doesn't exist"})
		return
	}

	ctx.Database.Debug().Model(&application).Update("name", params["name"])
	ctx.Response.JSON(&ApplicationResponse{Success: true, Application: &application})
}

// Delete will remove the resouce
// Route: DELETE /applications/{id}
func (c *ApplicationController) Delete(ctx *framework.Context) {
	params, valid := ctx.Request.Validate(map[string]string{
		"id": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

	var application models.Application
	ctx.Database.First(&application, params["id"])

	if application.ID == 0 {
		ctx.Response.JSON(&ApplicationResponse{Success: false, Message: "Application doesn't exist"})
		return
	}

	ctx.Database.Delete(&application)
	ctx.Response.JSON(&ApplicationResponse{Success: true})
}

// Test will render a home view
func (c *ApplicationController) Test(ctx *framework.Context) {

	v := &views.HomeView{}
	v.Init(ctx)

	ctx.Response.View(v)
}

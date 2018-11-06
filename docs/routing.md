# Routing

Our routing system works as a wrapper for the [Gorilla Mux Router](https://github.com/gorilla/mux). Routes live under the `routes` folder and here we have a couple of files you can use to define your routes.

* api.go: has the definition of our API routes
* web.go: has the definition of our web pages routes

?> Of course you can also define routes directly on the `main()` func but we recommend that you follow the files format to facilitate debugging and maintenance.

## Method Routes

To add a route we need to have the controller available in our project. Once we have it we can add routes using the request method names like follows:

```go
    c := &controllers.UsersController{}

    app.Router.Get("/users", c.Index)
    app.Router.Post("/users", c.Create)
    app.Router.Get("/users/{id}", c.Show)
    app.Router.Put("/users/{id}", c.Update)
    app.Router.Delete("/users/{id}", c.Delete)
```

?> Available methods here are `GET`, `POST`, `PUT` and `DELETE`.

## Resource Routes

As you see in the previous example, adding common routes for a resource can be a lot of work that will feel repetitive. To make this simpler we have the `Resource` method that will receive the full controller (implementing `IResourceController`) and we will add all those routes for you.

```go
    // This will do the same as the previous example
	app.Router.Resource("/users", &controllers.UserController{})
```

## Prefixes

It's pretty common to want to group routes under the same prefixed url when we are building APIs. This method will make every route refined in the closure have the prefix:

```go
	app.Router.Group("/api", func(r *framework.Router) {
		c := &controllers.UserController{}
		r.Get("/users", c.Index)
    })
```

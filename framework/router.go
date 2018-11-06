package framework

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router is a wrapper for the gorilla/mux router
type Router struct {
	*mux.Router
	middlewares []Middleware
	DB          *Database
}

// Context has all the important
// elements of our request flow
type Context struct {
	Request  *Request
	Response *Response
	Database *Database
}

// NewRouter will create a Router
func NewRouter() *Router {
	return &Router{
		&mux.Router{},
		make([]Middleware, 0),
		nil,
	}
}

func (r *Router) Use(mw Middleware) {
	r.middlewares = append(r.middlewares, mw)
}

func (r *Router) applyMiddlewares() {

}

// Get will add a route for a GET request method
func (r *Router) Get(path string, h func(*Context)) {
	r.makeRequest("GET", path, h)
}

// Post will add a route for a POST request method
func (r *Router) Post(path string, h func(*Context)) {
	r.makeRequest("POST", path, h)
}

// Put will add a route for a PUT request method
func (r *Router) Put(path string, h func(*Context)) {
	r.makeRequest("PUT", path, h)
}

// Delete will add a route for a Delete request method
func (r *Router) Delete(path string, h func(*Context)) {
	r.makeRequest("DELETE", path, h)
}

// Resource will add all routes we will need to add for a
// resource API. See the IResourceController
func (r *Router) Resource(path string, c IResourceController) {
	r.Get(path, c.Index)
	r.Get(path+"/{id}", c.Show)
	r.Post(path, c.Create)
	r.Put(path+"/{id}", c.Update)
	r.Delete(path+"/{id}", c.Delete)
}

// Group serves as a prefixer for a group or routes
func (r *Router) Group(prefix string, closure func(r *Router)) {
	prefixed := r.PathPrefix(prefix).Subrouter()
	router := &Router{prefixed, make([]Middleware, 0), r.DB}
	closure(router)
}

// makeRequest is the handler that we use to call all our
// routes. It prepares the context and calls the action
func (r *Router) makeRequest(
	method string,
	path string,
	handler func(*Context),
) {
	r.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {

		context := &Context{
			&Request{req},
			&Response{w},
			r.DB,
		}
		handler(context)

	}).Methods(method)
}

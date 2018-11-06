package framework

import (
	"encoding/json"
	"net/http"
)

// Response is a wrapper for the http.ResponseWriter
type Response struct {
	http.ResponseWriter
}

// String will print the passed string to the response
func (res *Response) String(s string) {
	res.Write([]byte(s))
}

// JSON will try to respond to the client a JSON Object
// we can use. Most common response in APIs for web
func (res *Response) JSON(el interface{}) {
	js, err := json.Marshal(el)
	if err != nil {
		res.String("Boooo")
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}

// View will use a template to render a view into the
// Response, it requires the layout and the template
func (res *Response) View(view IView) {

	templates, err := view.Render()
	if err != nil {
		http.Error(res, "500 Internal Server Error", 500)
		return
	}

	templates.ExecuteTemplate(res, view.GetLayout(), view.GetData())
}

# Context

The context is the struct instance that we send to our Controller actions. With this you have access to various elements in the request flow. Take a look at the definition:

```go
type Context struct {
	Request  *Request
	Response *Response
	Database *Database
}
```

## Request

The request is a wrapper for the internal `http.Request` so we have access to all of the functions available for it. Using it we have access to some helper functions to get data from our requests.

### Parameters

#### Query Parameters

This function will give you access to parameters that come from the query. An example of a parameter like this would be somehting like `http://example.com/example?foo=bar`.

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    value := req.QueryParam("foo")
    // do something with the query param value
}
```

#### URL Parameters

URL parameters are part of the route definition. When you have a router that has a path defined like this:

```go
app.Router.Get("/example/{id}", c.Show)
```

Then you can access that "id" parameters in your controller actions like this:

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    value = req.URLParam("id")
    // do something with the URL param value
}
```

#### JSON Parameters

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    value = req.JSONFormParam(key)
    // do something with the JSON param value
}
```

#### Form Parameters

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    value = req.FormParam(key)
    // do something with the Form param value
}
```

#### Any Parameter

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    value = req.Param(column)
    // do something with the Form param value
}
```

### Validations

To simplify parameter validations we have a function on the Request where you can define the validation rules for your parameters. Here's an example of how to use it:

```go
func (c *ApplicationController) Update(ctx *framework.Context) {

	params, valid := ctx.Request.Validate(map[string]string{
		"id":   "required",
		"name": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
    }

    // do something with params["id"] or params["name"]
}
```

#### Available validation rules

* required


## Response

A Response is a wrapper for the internal `http.ResponseWriter` so we have access to all of the functions available for it. Using it we have access to some helper functions to send data back to the client.

#### Print Strings

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    ctx.Response.String("This will be written to the response")
}
```

#### JSON

```go
type ApplicationResponse struct {
	Success    bool    `json:"success"`
	Data       *Data   `json:"data"`
}

func (c *ExampleController) Index(ctx *framework.Context) {
    var data Data
    ctx.Response.JSON(&ExampleResponse{Success: true, Data: &data})
}
```

This will return a json response like this:

```js
{
    "success": true,
    "data": "<your data>"
}
```

## Database

We initialize the database for you using the configurations you have on the `.env` file. This struct will give you access to all the GORM database functions.

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    var example models.Example
    ctx.Database.Debug().First(&example, 1)
    // do something with the example 1
}
```

?> For more information on how to use the `Database` and available functions please check out the official [GORM Documentation](http://gorm.io/docs/). There are also some examples on how this work on the Models guide.

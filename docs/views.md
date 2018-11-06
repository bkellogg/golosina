# Views

This guide will help you to render some HTML views from your controllers. Keep in mind that this part of the framework may change a bit.

Using the Response that we have in our controller action Context we are able to call a helper function to Render a view like this:

```go
func (c *ExampleController) Test(ctx *framework.Context) {

	v := &views.HomeView{}
	v.Init(ctx)

	ctx.Response.View(v)
}
```

A view has a couple of elements that we need to define so lets go step by step.

## Definition

The first thing we need to do is to create a file inside of the `views` folder with the following elements. The first element is the actual View definition, the second one is the data that we will send to it (like the view model):

```go
package views

type HomeView struct {
	Layout   string
	Template string
	Data     interface{}
}

type HomeViewData struct {
	Hello string
	World string
}
```

After this we need a some functions to be defined so we implement the IView interface:

#### Init

Here we will prepare all the data that our view will use. `Layout` is the name of the file inside the `templates/layouts` folder which will be our main template. The `Template` string is the name of the template we will use and the `Data` will be an instance of the ViewModel struct we defined before.

```go
func (v *HomeView) Init(ctx *framework.Context) {

	v.Layout = "app"
	v.Template = "home"

	v.Data = &HomeViewData{
		"Good",
		"Night",
	}
}
```

?> Keep in mind you have access to the context so you can prepare your data from the parameters.

#### Getters

Now lets define a couple of getters for the template paths:

```go
func (v *HomeView) GetLayout() string {
	return v.Layout
}

func (v *HomeView) GetData() interface{} {
	return v.Data
}
```

#### Render

Finally we need to add the code that will actually parse the templates into a usable object we can use on our Response. Here's an example of how we can code it:

```go
func (v *HomeView) Render() (*template.Template, error) {

	// prepare layout and template paths
	lp := path.Join("templates/layout", v.Layout+".html")
	fp := path.Join("templates", v.Template+".html")

	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("Template path issue")
		}
	}

	// Log if the file is a directory
	if info.IsDir() {
		return nil, errors.New("Template path is a directory")
	}

	// Parse the templates
	templates, err := template.ParseFiles(lp, fp)
	return templates, err

}
```


## Templates

On the templates folder you will have to add the layout and the template files. In this case I'm using files with the html format.

Let's begin with the layout (inside `views/templates/layout`) called `app.html`:

```html
{{define "app"}}
<!doctype html>
<html>
    <head>
        <meta charset="utf-8">
        <title>{{template "title"}}</title>
    </head>

    <body>
        {{template "body" .}}
    </body>
</html>
{{end}}
```

Ant now lets create in the `view/templates` folder a file called `home.html` with the following content:


```html
{{define "title"}}Home{{end}}

{{define "body"}}
    <h1>Welcome to Golosina</h1>
    <p>Hello: {{.Hello}}</p>
    <p>World: {{.World}}</p>
{{end}}
```

That's it, when you run the Test action that we defined at the beginning and try to request it you'll see the template rendered with out variables.

!> This whole view definition still feels too complex. We will work on simplifying it in the future

package views

import (
	"errors"
	"html/template"
	"os"
	"path"

	"github.com/eaperezc/golosina/framework"
)

// HomeView has view definition
type HomeView struct {
	Layout   string
	Template string
	Data     interface{}
}

// HomeViewData contains the data for this view
type HomeViewData struct {
	Hello string
	World string
}

// Init will prepare the data for the view
func (v *HomeView) Init(ctx *framework.Context) {

	v.Layout = "app"
	v.Template = "home"

	v.Data = &HomeViewData{
		"Good",
		"Night",
	}
}

// Render will take care of parsing the template
// so we can render it on our Response
func (v *HomeView) Render() (*template.Template, error) {

	// prepare layout and template paths
	lp := path.Join("views/templates/layout", v.Layout+".html")
	fp := path.Join("views/templates", v.Template+".html")

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

// GetLayout is a getter for the view Layout
func (v *HomeView) GetLayout() string {
	return v.Layout
}

// GetData is a getter for the view Data
func (v *HomeView) GetData() interface{} {
	return v.Data
}

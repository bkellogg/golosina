package framework

import (
	"html/template"
)

type IView interface {
	Init(*Context)
	Render() (*template.Template, error)
	GetLayout() string
	GetData() interface{}
}

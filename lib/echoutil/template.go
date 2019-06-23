package echoutil

import (
	"errors"
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// ITemplateRegistry ...
type ITemplateRegistry interface {
	Render(w io.Writer, name string, data interface{}, c echo.Context) error
}

// TemplateRegistry ...
type TemplateRegistry struct {
	Templates map[string]*template.Template
}

// NewEchoRenderer ...
func NewEchoRenderer(templateMap map[string]*template.Template) ITemplateRegistry {
	return &TemplateRegistry{Templates: templateMap}
}

// NewTemplateContainer ...
func NewTemplateContainer() map[string]*template.Template {
	return map[string]*template.Template{}
}

// RegisterTemplate ...
func RegisterTemplate(filenames ...string) *template.Template {
	var files []string
	for _, value := range filenames {
		files = append(files, fmt.Sprintf("template/%s.html", value))
	}

	// Templateで使用する関数を登録する
	funcMap := template.FuncMap{
		// "debugPrint": DebugPrint,
	}

	t := template.New(files[0]).Funcs(funcMap)

	return template.Must(t.ParseFiles(files...))
}

// Render ...
func (tr *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t, ok := tr.Templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}

	// Add global methods if data is a map
	if value, ok := data.(map[string]interface{}); ok {
		value["reverse"] = c.Echo().Reverse
	}

	return t.ExecuteTemplate(w, "layout", data)
}

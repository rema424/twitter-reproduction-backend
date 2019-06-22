package tmplutil

import (
	"github.com/flosch/pongo2"
)

const tmplPath = "static/dist/template/"

// HTML ...
func HTML(file string, data map[string]interface{}) (string, error) {
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).Execute(data)
}

// HTMLBlob ...
func HTMLBlob(file string, data map[string]interface{}) ([]byte, error) {
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}

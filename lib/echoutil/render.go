package echoutil

import (
	"net/http"

	"github.com/rema424/twitter-reproduction-backend/lib/tmplutil"

	"github.com/labstack/echo"
)

// Renderer ...
type Renderer struct {
	file403 string
	file404 string
	file500 string
	file502 string
}

// NewRenderer ...
func NewRenderer(file403 string, file404 string, file500 string, file502 string) Renderer {
	return Renderer{file403, file404, file500, file502}
}

// HTML ...
func (r Renderer) HTML(c echo.Context, code int, file string, data map[string]interface{}) error {
	if file == "" && code == http.StatusOK {
		return c.NoContent(http.StatusInternalServerError)
	}

	b, err := tmplutil.HTMLBlob(file, data)
	if err != nil {
		return c.NoContent(code)
	}

	return c.HTMLBlob(code, b)
}

// Render200 ...
func (r Renderer) Render200(c echo.Context, file string, data map[string]interface{}) error {
	b, err := tmplutil.HTMLBlob(file, data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.HTMLBlob(http.StatusOK, b)
}

// Render404 ...
func (r Renderer) Render404(c echo.Context, data map[string]interface{}) error {
	return r.HTML(c, http.StatusNotFound, r.file404, data)
}

// Render500 ...
func (r Renderer) Render500(c echo.Context, data map[string]interface{}) error {
	return r.HTML(c, http.StatusInternalServerError, r.file500, data)
}

// Render502 ...
func (r Renderer) Render502(c echo.Context, data map[string]interface{}) error {
	return r.HTML(c, http.StatusBadGateway, r.file502, data)
}

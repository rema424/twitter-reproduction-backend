package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// Response ...
type Response struct {
	Message string
}

// HelloHandler ...
func HelloHandler(c echo.Context) error {
	message := "Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!　Hello World!"

	// res := &Response{
	// 	Message: message,
	// }

	// return c.JSON(http.StatusOK, res)

	data := map[string]interface{}{"Message": message, "Number": 2364821976}

	// b, err := tmplutil.HTMLBlob("example.html", data)
	// if err != nil {
	// 	return err
	// }
	// return c.HTMLBlob(http.StatusOK, b)

	// s, err := tmplutil.HTML("example.html", data)
	// if err != nil {
	// 	return err
	// }
	// return c.HTML(http.StatusOK, s)

	// return renderer.Render502(c, data)
	// return renderer.HTML(c, http.StatusOK, "example.html", data)
	return renderer.HTML(c, http.StatusOK, "", data)
	// return echoutil.Render200(c, "example.com", data)
}

// ParrotHandler ...
func ParrotHandler(c echo.Context) error {
	message := c.Param("message")

	res := &Response{
		Message: message,
	}

	return c.JSON(http.StatusOK, res)
}

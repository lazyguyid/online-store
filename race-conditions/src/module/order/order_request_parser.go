package order

import "github.com/labstack/echo"

// EchoBinder func
func (request *Request) EchoBinder(c echo.Context) (err error) {
	if err = c.Bind(request); err != nil {
		return
	}

	return nil
}

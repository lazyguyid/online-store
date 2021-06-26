package core

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// EchoStart func
func EchoStart(app App) error {
	e := echo.New()
	cfg := app.Config()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	fmt.Println("\nRegister Path...")
	fmt.Println("======================================")
	for _, uc := range app.Usecases() {
		path := fmt.Sprintf("%s/%s", uc.Version, uc.RestPath)
		h := newHandler(uc)

		fmt.Println("⇨ Registering Path: ", path)
		switch uc.HTTPMethod {
		case "GET":
			e.GET(path, h.Handler, uc.Middleware.Echo)
			break
		case "POST":
			e.POST(path, h.Handler, uc.Middleware.Echo)
			break
		case "PUT":
			e.PUT(path, h.Handler, uc.Middleware.Echo)
			break
		case "DELETE":
			e.DELETE(path, h.Handler, uc.Middleware.Echo)
			break
		}
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Get("PORT").(string))))

	return nil
}

type handler struct {
	uc *ConfigUsecase
}

func newHandler(uc *ConfigUsecase) *handler {
	h := new(handler)
	h.uc = uc

	return h
}

func (h *handler) Handler(c echo.Context) error {

	r := h.uc.RequestParser()
	err := r.EchoBinder(c)
	if err != nil {
		return err
	}

	result := h.uc.Usecase(NewContext(ConfigContext{Echo: c}), r.GetInstance())

	res := make(map[string]interface{})

	res["message"] = result.Message
	res["code"] = result.Code
	res["data"] = result.Data

	if result.Error != nil {
		res["error"] = result.Error.Error()
		// return result.Error
	}

	return c.JSON(result.Code, res)
}

package main

import (
	"fmt"
	"os"

	"online-store/config"
	"online-store/core"
	"online-store/deps"
	"online-store/src/module/order"
	"online-store/src/shared/helpers"
	"online-store/src/shared/repositories"

	"github.com/joho/godotenv"
)

func main() {

	// load ENV
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	// load config
	app := core.Application(config.Load())
	app.SetStorage(config.NewStorage(app))

	// register helper
	app.RegisterHelper(helpers.NewHelper(app))

	// register repositories
	app.AddCRepository("order", repositories.NewOrderRepository(app))
	app.AddCRepository("product", repositories.NewProductRepository(app))
	// app.AddCRepository("user", repositories.NewUserRepository(app))

	oc := order.NewOrderUsecase(app)

	app.Register(&core.ConfigUsecase{
		RestPath:      "cart/buy",
		HTTPMethod:    "POST",
		Usecase:       oc.Buy,
		RequestParser: order.NewOrderRequest,
		Middleware: &core.Middleware{
			Echo: deps.EchoAuthMiddleware,
		},
		Enable: true,
	})

	app.Start()
}

package main

import (
	"fmt"
	"os"

	"online-store/config"
	"online-store/deps"
	"online-store/src/module/order"
	"online-store/src/shared/helpers"
	"online-store/src/shared/repositories"

	"github.com/joho/godotenv"
	"github.com/lazyguyid/gacor"
)

func main() {

	// load ENV
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	// load config
	app := gacor.Application(config.Load())
	app.SetStorage(config.NewStorage(app))

	// register helper
	app.RegisterHelper(helpers.NewHelper(app))

	// register repositories
	app.AddCRepository("order", repositories.NewOrderRepository(app))
	app.AddCRepository("product", repositories.NewProductRepository(app))
	// app.AddCRepository("user", repositories.NewUserRepository(app))

	oc := order.NewOrderUsecase(app)

	app.Register(&gacor.ConfigUsecase{
		RestPath:      "cart/buy",
		HTTPMethod:    "POST",
		Usecase:       oc.Buy,
		RequestParser: order.NewOrderRequest,
		Middleware: &gacor.Middleware{
			Echo: deps.EchoAuthMiddleware,
		},
		Enable: true,
	})

	app.Start()
}

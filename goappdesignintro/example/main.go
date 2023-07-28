package main

import (
	"log"
	"net/http"

	"github.com/bagardavidyanisntreal/go-lessons/goappdesignintro/example/app"
	"github.com/bagardavidyanisntreal/go-lessons/goappdesignintro/example/order"
	"github.com/bagardavidyanisntreal/go-lessons/goappdesignintro/example/repository/imdb"
	"github.com/bagardavidyanisntreal/go-lessons/goappdesignintro/example/repository/pgsql"
)

func main() {
	pgsqlRepository := pgsql.NewRepository()
	imdbOrders := imdb.NewOrder(pgsqlRepository)

	// here we can choose which repo to use: pgsql or imdb, or even any other
	orderService := order.NewService(imdbOrders)
	application := app.NewApplication(orderService)

	mux := http.NewServeMux()
	mux.HandleFunc("/order/status", application.UpdateOrderStatus)

	if err := http.ListenAndServe("", mux); err != nil {
		log.Fatal(err)
	}
}

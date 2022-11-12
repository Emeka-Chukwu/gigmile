package main

import (
	"gigmile-task/config"
	"gigmile-task/data"
	"gigmile-task/handler"
	"gigmile-task/routes"
	"log"
	"net/http"
	"os"

	_ "gigmile-task/routes"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

func main() {

	conn := config.ConnectToDB()
	if conn == nil {
		log.Panic("Can't connect to postgres!")
	}
	app := handler.HandlerModel{
		DB:     conn,
		Models: data.New(conn),
	}

	r := mux.NewRouter()

	routes.CountryRouter(r, app)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+webPort, r)
}

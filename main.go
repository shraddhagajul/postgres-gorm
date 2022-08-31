package main

import (
	"log"
	"net/http"
	"postgres/routes"
	"postgres/services"
	"postgres/utility"
)

func main() {
	db := utility.GetConnection()
	services.SetDB(db)
	appRouter := routes.CreateRoutes()

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", appRouter))
}

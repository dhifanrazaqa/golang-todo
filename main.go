package main

import (
	"github.com/dhifanrazaqa/golang-todo/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	port, exist := os.LookupEnv("PORT")

	if !exist {
		log.Fatal("PORT not set in .env")
	}

	err := http.ListenAndServe("localhost:"+port, routes.Init())

	if err != nil {
		log.Fatal(err)
	}
}
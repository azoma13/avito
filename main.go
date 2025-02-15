package main

import (
	"log"
	"net/http"
	"os"

	"github.com/azoma13/avito/configs"
	"github.com/azoma13/avito/internal/dataBase"
	"github.com/azoma13/avito/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {

	conn := dataBase.ConnectToDB()
	defer conn.Close()

	port := os.Getenv("PORT")

	if port == "" {
		port = configs.DefaultPort
	}
	port = ":" + port
	r := mux.NewRouter()
	r.HandleFunc("/api/auth", handlers.AuthEmployeeHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/buy/{item}", handlers.BuyItemHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/sendCoin", handlers.SendCoinHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/info", handlers.InfoEmployeeHandler).Methods(http.MethodPost)

	log.Println("application running on port" + port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		panic(err)
	}
}

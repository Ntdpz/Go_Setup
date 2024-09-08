package main

import (
	"fmt"
	"go-backend-api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HelloGolang).Methods("GET")
	r.HandleFunc("/api/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/api/projects", handlers.GetProjects).Methods("GET")

	// เริ่มเซิร์ฟเวอร์
	fmt.Println("Server started at :7777")
	log.Fatal(http.ListenAndServe(":7777", r))
}

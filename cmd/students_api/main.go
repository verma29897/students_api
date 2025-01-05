package main

import (
	"fmt"
	"net/http"

	"github.com/verma29897/students_api/internal/config"
)

func main() {
	fmt.Println("hello First App")
	//load config
	cfg := config.MustLoad()
	//database setup
	//setup router
	router := http.NewServeMux()
	router.HandleFunc("GET/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello, world!"))
	})

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Println("server started")

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to server", err)
	}
}

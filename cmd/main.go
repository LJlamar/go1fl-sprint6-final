package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	//ЛОГГЕР
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	logger.Println("Starting server at port 8080")

	s := server.NewServer(logger)
	err := http.ListenAndServe(s.Server.Addr, s.Server.Handler)
	if err != nil {
		logger.Fatalln("server failed: ", err)
	}
}

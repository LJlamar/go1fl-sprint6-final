package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Application struct {
	Server *http.Server
	Logger *log.Logger
}

func NewServer(logger *log.Logger) *Application {

	//logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	// "Регистрируем два новых обработчика и соответствующие URL-шаблоны в
	// маршрутизаторе servemux"
	//HTTP-роутер:
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	app := Application{
		Server: srv,
		Logger: logger,
	}

	return &app
	/*
	   fmt.Println("Starting server at port 8080")
	   err := Srv.ListenAndServe()

	   	if err != nil {
	   		fmt.Println("Error starting the server:", err)
	   	}
	*/
}

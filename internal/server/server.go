package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server структура, инкапсулирующая логгер и http-сервер
type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

// NewServer инициализирует и возвращает сервер с зарегистрированными маршрутами
func NewServer(logger *log.Logger) *Server {
	// 1. Создание маршрутизатора
	mux := http.NewServeMux()

	// 2. Регистрация хендлеров
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	// 3. Настройка http.Server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		HTTP:   srv,
	}
}

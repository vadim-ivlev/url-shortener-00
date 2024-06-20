package main

import (
	"net/http"
)

func main() {

	// Создать мультиплексор
	mux := http.NewServeMux()

	// Добавить обработчик для корневого URL
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request). {
		method := r.Method
		// Эндпоинт с методом `POST`
		if method == "POST" {
			shortenedURL := getShortenedURL(r)

		w.Write([]byte("Hello, World! " + method + "\n"))
	})

	// Запустить сервер на порту 8080
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}

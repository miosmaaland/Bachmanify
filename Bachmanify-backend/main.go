package main

import (
	"ErlichBachmanify-backend/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", handlers.SocketHandler)
	http.HandleFunc("/progress", handlers.GetProgress)
	http.HandleFunc("/progress/save", handlers.SaveProgress)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", handlers.SocketHandler)

	println("🚀 ErlichBachmanify backend running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

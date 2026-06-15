package main

import (
	"awesomeLearnigGO/ToDoList"
	"awesomeLearnigGO/ToDoList/http"
	"log"
)

func main() {
	todoList := ToDoList.NewList()

	handlers := http.NewHTTPHandlers(todoList)

	server := http.NewHTTPServer(handlers)

	log.Println("[INFO] Server is starting on http://localhost:9091...")
	if err := server.StartServer(); err != nil {
		log.Fatalf("[FATAL] Server failed: %v", err)
	}
}

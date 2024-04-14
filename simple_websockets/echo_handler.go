package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var mutex = sync.Mutex{}

func websocketEndpoint(writer http.ResponseWriter, request *http.Request) {

	connection, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer connection.Close()

	mutex.Lock()
	clients[connection] = true
	mutex.Unlock()
	log.Println("Client connected")
	defer func() {
		log.Println("Client disconnected")
		mutex.Lock()
		delete(clients, connection)
		mutex.Unlock()
	}()
	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("Received message from client: %s\n", message)
		broadcast(messageType, message)
	}
}
func homepage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the Schwarf's webSocket chat server!")
}

func broadcast(messageType int, message []byte) {
	mutex.Lock()
	defer mutex.Unlock()
	for client := range clients {
		if err := client.WriteMessage(messageType, message); err != nil {
			log.Printf("Error wrting to WebSocket: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/ws", websocketEndpoint)
}

func main() {
	// http.HandleFunc("/echo", echoHandler)
	// fmt.Println("Running server ... ")
	// http.ListenAndServe(":8080", nil)
	fmt.Println("Hello")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

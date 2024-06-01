package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Client struct {
	ID         string
	Connection *websocket.Conn
}

var clients = make(map[*Client]bool)
var mutex = sync.Mutex{}
var idCounter = 0

func websocketEndpoint(writer http.ResponseWriter, request *http.Request) {

	connection, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer connection.Close()

	idCounter++
	clientID := fmt.Sprintf("Client-%d", idCounter)
	client := &Client{ID: clientID, Connection: connection}

	mutex.Lock()
	clients[client] = true
	mutex.Unlock()
	log.Printf("Client %s connected", client.ID)
	defer func() {
		log.Printf("Client %s disconnected", client.ID)
		mutex.Lock()
		delete(clients, client)
		mutex.Unlock()
	}()
	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("Received message from client %s at s%: %s\n", client.ID, time.Now().Format(time.RFC3339), message)
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
		if err := client.Connection.WriteMessage(messageType, message); err != nil {
			log.Printf("Error wrting to WebSocket: %v", err)
			client.Connection.Close()
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

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
	log.Println("Client connected")

	err = connection.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	reader(connection)
}

func reader(connection *websocket.Conn) {
	for {
		messgaeType, p, err := connection.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))

		if err := connection.WriteMessage(messgaeType, p); err != nil {
			log.Println(err)
			return
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

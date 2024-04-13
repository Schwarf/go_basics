package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// echoHandler handles WebSocket requests from the peer.
// func echoHandler(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Println("Error upgrading to WebSocket:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	for {
// 		mt, message, err := conn.ReadMessage()
// 		if err != nil {
// 			fmt.Println("Error reading message:", err)
// 			break
// 		}
// 		fmt.Printf("Received message: %s\n", message)
// 		err = conn.WriteMessage(mt, message)
// 		if err != nil {
// 			fmt.Println("Error writing message:", err)
// 			break
// 		}
// 	}
// }

func homepage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Homepage")

}

func websocketEndpoint(writer http.ResponseWriter, request *http.Request) {
	upgrader.CheckOrigin = func(request *http.Request) bool { return true }
	connection, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer connection.Close()
	log.Println("Client connected")
	// for {
	// 	_, message, err := connection.ReadMessage()
	// 	if err != nil {
	// 		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
	// 			log.Printf("error: %v", err)
	// 		}
	// 		break
	// 	}
	// 	log.Printf("Received: %s", message)

	// 	// Echo the message back
	// 	if err := connection.WriteMessage(websocket.TextMessage, message); err != nil {
	// 		log.Println("write:", err)
	// 		break
	// 	}
	// }

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

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Schwarf/go_basics/simple_websockets/encryption"

	create_db "github.com/Schwarf/go_basics/simple_websockets/sql"

	"github.com/gorilla/websocket"
)

func storeMessage(clientID string, message []byte, timestamp string, key []byte) error {
	plainText := fmt.Sprintf("Client: %s, Timestamp: %s, Message: %s\n", clientID, timestamp, message)
	encryptedMessage, err := encryption.Encrypt(plainText, key)
	decryptedMessage, err := encryption.Decrypt(encryptedMessage, key)
	log.Printf("Decrypted message: %s", decryptedMessage)
	if err != nil {
		log.Printf("Encryption did not work! Error: %v", err)
		return err
	}

	file, err := os.OpenFile("messages.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Opening file did not work! Error: %v", err)
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(encryptedMessage + "\n"); err != nil {
		log.Printf("Writing string did not work! Error: %v", err)
		return err
	}

	return nil
}

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

	encryptionKey := []byte("your-32-byte-long-key-here123456")

	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		timestamp := time.Now().Format(time.RFC3339)
		log.Printf("Received message from client %s at s%: %s\n", client.ID, time.Now().Format(time.RFC3339), message)
		if err := storeMessage(clientID, message, timestamp, encryptionKey); err != nil {
			log.Printf("Failed to store message! Error: %v", err)
		}
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
	db, err := create_db.ConnectToDatabase()
	if err != nil {
		log.Printf("Database connection failed! %v", err)
	}
	err = create_db.CreateMessagesTable(db)
	if err != nil {
		log.Printf("Creation of table failed! %v", err)
	}

	fmt.Println("Hello")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

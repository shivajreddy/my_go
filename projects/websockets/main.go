package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// HUB
type Hub struct {
	Clients map[*Client]bool

	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message

	Mutex sync.Mutex
}

func (hub *Hub) run() {
	for {
		select {

		case client := <-hub.Register:
			hub.Mutex.Lock()
			hub.Clients[client] = true
			hub.Mutex.Unlock()

		case client := <-hub.Unregister:
			hub.Mutex.Lock()
			delete(hub.Clients, client)
			close(client.Send) // critical
			hub.Mutex.Unlock()

		// send the message to send chanel for all clients
		case message := <-hub.Broadcast:
			hub.Mutex.Lock()
			for client := range hub.Clients {
				select {
				case client.Send <- message:
				default:
					delete(hub.Clients, client)
					close(client.Send)
				}
			}
			hub.Mutex.Unlock()
		}
	}
}

// Message
type Message struct {
	ClientId string
	Data     string
}

// Client
type Client struct {
	Id string

	Conn *websocket.Conn
	Send chan Message

	Hub *Hub
}

func (client *Client) readPump() {
	defer func() {
		client.Hub.Unregister <- client
		client.Conn.Close()
	}()

	for {
		_, data, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("failed to read")
			return
		}
		message := Message{ClientId: client.Id, Data: string(data)}
		client.Hub.Broadcast <- message
	}
}

func (client *Client) writePump() {
	defer client.Conn.Close()
	for message := range client.Send {
		if err := client.Conn.WriteJSON(message); err != nil {
			log.Println("Couldn't send")
			break
		}
	}
}

func serveWebsocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	w.Header().Set("Content-Type", "application/json")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("failed to upgrade http to websocket")
		return
	}

	client := Client{
		Id:   uuid.New().String(),
		Hub:  hub,
		Conn: conn,
		Send: make(chan Message),
	}

	hub.Register <- &client

	go client.readPump()
	go client.writePump()
}

func main() {
	fmt.Println("WEBSOCKETS")

	hub := Hub{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
	}
	go hub.run()
	log.Println("started hub")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(&hub, w, r)
	})

	// Start HTTP server
	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}

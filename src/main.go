package main

import (
  "log"
  "net/http"
  "github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // map of connected clients
var broadcast = make(chan Message) // broadcast channel
var upgrader = websocket.Upgrader{}

// Define our message object
type Message struct {
  Email string `json:"email"`
  Username string `json:"username"`
  Message string `json:"message"`
}

func main() {
  // File server
  fs := http.FileServer(http.Dir("./public"))
  http.Handle("/", fs)

  // Configure websocket route
  http.HandleFunc("/ws", handleConnections)

  // Start a goroutine listening for messages
  go handleMessages()

  log.Println("HTTP Server listening on :8181")
  err := http.ListenAndServe(":8181", nil)
  if err != nil {
    log.Fatal("HTTP Server failed: ", err)
  }
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
  // Upgrade initial GET request into a websocket
  ws, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Fatal("WS", err)
  }

  // Make sure we close the connection when the function returns
  defer ws.Close()

  // Register our new client
  clients[ws] = true

  // Loop, waiting for a reply
  for {
    var msg Message

    // Read the message and map it to our Message Struct
    err := ws.ReadJSON(&msg)

    // If there was an error of any sort we can assume that the client is no longer connected, and remove
    // them from the map
    if err != nil {
      log.Printf("Error: %v", err)
      delete(clients, ws)
      break
    }

    // Some black magic to send the received message to the broadcaster
    broadcast <- msg
  }
}

func handleMessages() {
  // Loop, broadcasting messages
  for {
    // Grab the next messaage from the broadcast channel
    msg := <-broadcast
    // Send it out to every client that is currently connected
    for client := range clients {
      err := client.WriteJSON(msg)
      // If there was an error broadcasting to a client, we can assume that the client is no longer connected
      // adn remove them from the map
      if err != nil {
        log.Printf("Failed to broadcast: %v", err)
        client.Close()
        delete(clients, client)
      }
    }
  }
}

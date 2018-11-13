package main

import (
  "log"
  "net/http"
  "time"
  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
  // Display an index.html
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
  })

  // Receive a WebSocket message, relay it back to the client
  http.HandleFunc("/v1/ws", func (w http.ResponseWriter, r *http.Request) {
    var conn, _ = upgrader.Upgrade(w, r, nil) // third param is headers go func(conn *websocket.Conn) {
    go func(conn *websocket.Conn) {
      for {
        mType, msg, _ := conn.ReadMessage()

        conn.WriteMessage(mType, msg)
      }
    }(conn)
  })

  // Receive a WebSocket message, log it to the server
  http.HandleFunc("/v2/ws", func (w http.ResponseWriter, r *http.Request) {
    var conn, _ = upgrader.Upgrade(w, r, nil)
    go func(conn *websocket.Conn) {
      for {
        _, msg, _ := conn.ReadMessage()

        log.Print("v2 socket response: ", string(msg))
      }
    }(conn)
  })

  // Send a message over a WebSocket every 5 seconds
  http.HandleFunc("/v3/ws", func (w http.ResponseWriter, r *http.Request) {
    var conn, _ = upgrader.Upgrade(w, r, nil)
    go func(conn *websocket.Conn) {
      ch := time.Tick(5 * time.Second)
      for range ch {
        conn.WriteJSON(MessageStruct{
          Username: "lukeberry99",
          FirstName: "Luke",
          LastName: "Berry",
        })
      }
    }(conn)
  })

  // Handle the closing of a WebSocket
  http.HandleFunc("/v4/ws", func (w http.ResponseWriter, r *http.Request) {
    var conn, _ = upgrader.Upgrade(w, r, nil)
    go func(conn *websocket.Conn) {
      for {
        _, _, err := conn.ReadMessage()

        if err != nil {
          conn.Close()
        }
      }
    }(conn)
  })

  log.Print("HTTP Server listening on :3000")
  err := http.ListenAndServe(":3000", nil)
  if err != nil {
    log.Fatal("HTTP Server crashed: %v", err)
  }
}

type MessageStruct struct {
  Username string `json:"username"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
}

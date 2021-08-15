package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Msg  string `json:"msg,omitempty"`
	Name string `json:"name,omitempty"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketConnectHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
		return
	}
	clients[conn] = true
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	var msg Message
	msg.Msg = r.FormValue("msg")
	msg.Name = r.FormValue("name")
	broadcast <- msg
}

func websocketMessage() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	portNumber := "9000"
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/ws", websocketConnectHandler)
	http.HandleFunc("/msg", messageHandler)
	log.Println("Server listening on port", portNumber)
	go websocketMessage()
	http.ListenAndServe(":"+portNumber, nil)
}

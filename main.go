package main

import (
	"chatting_back/pkg/websocket"
	"fmt"
	"log"
	"net/http"
)

func serveWS(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWS)
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

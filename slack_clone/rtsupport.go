package main // creates an executable

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	// pass in URL pattern and handler fcn to execute
	http.HandleFunc("/", handler)
	// waits until port is connected to or app is blocked
	http.ListenAndServe(":4000", nil)

}

// upgrader "upgrades" the HTTP server connection to the WebSocket protocol.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// returning true allows WebSocket connection from any origin
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	// pass in a writer and string we'd like written as a response
	// fmt.Fprintf(w, "this is how we create a basic HTTP web server in Go")

	var socket, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// if upgrade is successful, read and write message
	for {
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(msg))
		// write msg back to the client
		// an 'if' condition allows you to specify statement before conditional
		if err := socket.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err)
			return
		}
	}
}

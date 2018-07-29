package main // creates an executable

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Message contains name and data. Fields need to be capitalized to be made public.
// json.Marshal is not part of the main package, so it doesn't have access to
// fields in lowercase.
type Message struct {
	Name string `json:"name"` // add field tags, a kind of metadata to specify
	// special encoding
	Data interface{} `json:"data"` // empty interface defines no methods, which means that every
	// data type in Go implements the behavior of an empty interface.
}

// Channel holds name and ID
type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	// pass in URL pattern and handler fcn to execute
	http.HandleFunc("/", handler)
	// waits until port is connected to or app is blocked
	fmt.Println("Listening on port 4000...")
	http.ListenAndServe(":4000", nil)
}

// upgrader "upgrades" the HTTP server connection to the WebSocket protocol.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// modern web browsers allow web socket connections to any host. We want a
	// no browser based same-origin policy. Gorilla websockets
	// default behavior forces a same origin policy. Lets leave origin policy for
	// the server to decide. The server will determine if it's okay for the browser
	// to make a websocket connection.
	//
	// Assign a new fcn to CheckOrigin and return TRUE indicating we'll allow
	// connections from any origin.
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	// pass in a writer and string we'd like written as a response
	// fmt.Fprintf(w, "this is how we create a basic HTTP web server in Go")
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// if Upgrade is successful, read and write message in websocket
	for {
		// ReadMessage is a fcn that blocks until a message is received.
		// msgType, msg, err := socket.ReadMessage()
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// fmt.Println(string(msg), msgType)

		var inMsg Message
		if err := socket.ReadJSON(&inMsg); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%#v\n", inMsg) // adds data type for field
		fmt.Printf("%v\n", inMsg)
		// echo the msg back to the client
		// an 'if' condition allows you to specify statement before conditional
		// if err := socket.WriteMessage(msgType, msg); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
	}
}

package main

import (
	"fmt"

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

// Client is responsible for reading and writing to the web socket. So Client will
// need to have access to the websocket.
type Client struct {
	send   chan Message
	socket *websocket.Conn
}

func (client *Client) Read() {
	var msg Message
	for {
		if err := client.socket.ReadJSON(&msg); err != nil {
			break
		}
	}
}

// method of Client to send messages to the client over the websocket.
func (client *Client) Write() {
	for msg := range client.send {
		fmt.Printf("%#v\n", msg)
		if err := client.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	client.socket.Close()
}

func NewClient(socket *websocket.Conn) *Client {
	return &Client{
		send:   make(chan Message),
		socket: socket,
	}
}

// Example:
// func main() {
// 	msgChan := make(chan string)
// 	// anonymous fcn that's immediately invoked. "Hello" was sent in a separate
// 	// go routine and was sent to the original go routine that's running main fcn.
// 	go func() {
// 		msgChan <- "Hello"
// 	}()

// 	msg := <-msgChan
// 	fmt.Println(msg)
// }

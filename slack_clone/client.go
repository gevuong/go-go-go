package main

import (
	"fmt"
	"math/rand"
	"time"
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

// Client stores send channel
type Client struct {
	send chan Message
}

// method of Client to send messages to the client over the websocket.
func (client *Client) write() {
	for msg := range client.send {
		fmt.Printf("%#v\n", msg)
	}
}

// method to start DB query that will stream channel changes (i.e. add, edit, delete)
func (client *Client) subscribeChannels() {
	for {
		// provides a random time duration of up to 1sec.
		time.Sleep(r())
		client.send <- Message{"channel add", ""}
	}
}

// method to start DB query that will stream channel changes (i.e. add, edit, delete)
func (client *Client) subscribeMessages() {
	for {
		// provides a random time duration of up to 1sec.
		time.Sleep(r())
		client.send <- Message{"message add", ""}
	}
}

// a throwaway fcn that returns a random duration between 0-1 sec.
func r() time.Duration {
	return time.Millisecond * time.Duration(rand.Intn(1000))
}

func NewClient() *Client {
	return &Client{
		send: make(chan Message),
	}
}
func main() {
	client := NewClient()
	go client.subscribeChannels()
	go client.subscribeMessages()
	client.write()

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

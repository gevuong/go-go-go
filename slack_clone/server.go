package main // creates an executable

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
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
	router := NewRouter()
	// When router receives "channel add" msg, call the addChannel fcn.
	router.Handle("channel add", addChannel)

	// pass in URL pattern and handler fcn to execute
	http.Handle("/", router)
	// waits until port is connected to or app is blocked
	fmt.Println("Listening on port 4000...")
	http.ListenAndServe(":4000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

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
		var outMsg Message
		if err := socket.ReadJSON(&inMsg); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%#v\n", inMsg) // adds data type for field when logging
		fmt.Printf("%v\n", inMsg)

		switch inMsg.Name {
		case "channel add":
			err = addChannel(inMsg.Data)
			if err != nil {
				outMsg = Message{"error", err}
				if err := socket.WriteJSON(outMsg); err != nil {
					fmt.Println(err)
					break
				}
			}
		case "channel subscribe":
			go subscribeChannel(socket) // take a socket parameter
		}
		// echo the msg back to the client
		// an 'if' condition allows you to specify statement before conditional
		// if err := socket.WriteMessage(msgType, msg); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
	}
}

func addChannel(data interface{}) error {
	var channel Channel
	// use mapstructure pkg for type checking and value checking, otherwise panic will
	// occur if assertions are incorrect.
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}
	// type assertion tells Go that data and name contains a more specific type.
	// For example, value of name property is a string.
	// I guess its similar to PropTypes in React.
	// channelMap := data.(map[string]interface{})
	// channel.Name = channelMap["name"].(string)

	channel.ID = "1" // ID is normally set by rethinkDB when record is inserted.
	// fmt.Printf("%v\n", channel)
	fmt.Println("added channel")
	return nil
}

// take a socket parameter that is a pointer to a websocket connection
func subscribeChannel(socket *websocket.Conn) {
	// rethinkDB Query / changefeed
	for {
		// pause for one second in for loop
		time.Sleep(time.Second * 1)
		// construct a new message and encode to JSON
		msg := Message{"channel add",
			Channel{"1", "software support"}}
		socket.WriteJSON(msg)
		fmt.Println("sent new channel")
	}
}

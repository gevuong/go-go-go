package main

import (
	"encoding/json"
	"fmt"

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

// An interface specifies the behavior of an object. For example, the behavior of
// a speaker interface is the ability to speak. Interfaces are a way of generalizing.
type speaker interface {
	Speak()
}

// Message struct now has a speak method, which is a fcn that has access to
// the fields in the Message struct, which is pretty cool. In other words, Message
// struct implements the behavior of a speaker by implementing the speak method.
func (m Message) speak() {
	fmt.Println("I'm a " + m.Name + " event!")
}

// we can also write a fcn that takes a speaker parameter, which is basically saying,
// if you can speak, you can be used here, or use this fcn I believe.
func someFunc(speaker speaker) {
	speaker.Speak()
}

func main() {
	recRawMsg := []byte(
		`{"name": "channel add",` +
			`"data":{"name":"Hardware Support"}}`)

	var recMsg Message
	// Unmarshal is expecting a pointer, so we can pass pointer to recMsg
	if err := json.Unmarshal(recRawMsg, &recMsg); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", recMsg)

	if recMsg.Name == "channel add" {
		channel, err := addChannel(recMsg.Data) // pass in channel info

		// now lets assume that the channel was successfully added to the database,
		// and lets send message across the websocket
		var sendMsg Message
		sendMsg.Name = "channel add"
		sendMsg.Data = channel

		sendRawMsg, err := json.Marshal(sendMsg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(sendRawMsg))
	}
}

func addChannel(data interface{}) (Channel, error) {
	var channel Channel
	// use mapstructure pkg for type checking and value checking, otherwise panic will
	// occur if assertions are incorrect.
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return channel, err
	}
	// type assertion tells Go that data and name contains a more specific type.
	// For example, value of name property is a string.
	// I guess its similar to PropTypes in React.
	// channelMap := data.(map[string]interface{})
	// channel.Name = channelMap["name"].(string)
	channel.ID = "1" // ID is normally set by rethinkDB when record is inserted.
	fmt.Printf("%v\n", channel)
	return channel, nil
}

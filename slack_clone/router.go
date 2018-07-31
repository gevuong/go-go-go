package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Handler is of type func()
type Handler func(*Client, interface{})

// Router is empty
type Router struct {
	rules map[string]Handler // remember, maps need to be initialized before they're used
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

// NewRouter creates a new router
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

// Handle
func (e *Router) Handle(msgName string, handler Handler) {
	e.rules[msgName] = handler
}

func (e *Router) serveHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "this is how we create a basic HTTP web server in Go")
	// Upgrade upgrades HTTP request connection to a websocket connection
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	client := NewClient(socket)
	// methods below will be running independently, so they need to be in separate
	// go routines
	go client.Write()
	client.Read()
}

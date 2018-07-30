# README

- We will only be using Go for websocket traffic. 
- We will serve frontend HTML and JS to dedicated servers such as Nginx, or CDN if appropriate. Leave normal HTTP traffic to Nginx or CDN.
- Websockets allow connections to any host. There are no browser-based same origin policy. The origin policy is left to the server to decide. In other words, it's up to the server to determine if it's okay for the web browser to make a websocket connection.

- To test websocket communication: `http://websocket.org/echo.html`. Input whichever localhost you are using (i.e. `ws://localhost:4000`), click "CONNECT", and make sure log says "CONNECTED". Then input a message to send. And you'll see the message sent and echoed back to us. You will also see the message logged in the terminal.


- To test `socket.ReadJSON` and `socket.WriteJSON`, go to `jsbin.com` and add following JS code
    ```
    let msg = {
    name: "channel add",
    data: {
        name: "hardware support"
    }
    }

    let ws = new WebSocket('ws://localhost:4000');

    ws.onopen = () => {
    ws.send(JSON.stringify(msg));
    }
    ```
- Next, go run main.go and click Run JS code in `jsbin`. Terminal should log message struct.


### Goroutines 
- Think of it as a separate execution thread. This means that if a fcn (i.e. subscribeChannel) is originally blocking, and we still want to receive new messages in the meantime, adding a go routine to fcn not block the flow control of the main program. 
- Very easy to create, by simply adding `go` keyword in front of a function call. 
- Very lightweight in comparison to other languages. A goroutine requires very little server resources, a couple kB of stack memory. **This means you can literally create thousands of go routines without crashing your server.** While other languages require a couple orders of magnitude more memory for each thread created.
- As a result, go routines make creating highly concurrent applications fairly simple.
- Go routines allow us to run multiple separate tasks, concurrently.
**Must avoid situations where two or more goroutines are modifying shared data at the same time.** This can lead to race conditions.
- Channels provide a safe way of passing values between go routines. There is no risk of having the go routines accessing whatever values passed at the same time.
- Popular quote in Go: "Do not communicate by sharing memory, share memory by communicating."

- Run on jsbin to test subscribing and adding channel concurrently using go routines.
    ```
    let msg = {
    name: 'channel add',
    data: {
        name: 'hardware support'
    }
    }

    let subMsg = {
    name: 'channel subscribe'
    }

    let ws = new WebSocket('ws://localhost:4000');

    ws.onopen = () => {
    ws.send(JSON.stringify(subMsg))
    ws.send(JSON.stringify(msg));
    }

    ws.onmessage = (e) => {
    console.log(e.data);
    }
    ```
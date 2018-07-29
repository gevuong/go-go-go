# README

- We will only be using Go for websocket traffic. 
- We will serve frontend HTML and JS to dedicated servers such as Nginx, or CDN if appropriate. Leave normal HTTP traffic to Nginx or CDN.
- Websockets allow connections to any host. There are no browser-based same origin policy. The origin policy is left to the server to decide. In other words, it's up to the server to determine if it's okay for the web browser to make a websocket connection.

- To test websocket communication: `http://websocket.org/echo.html`. Input whichever localhost you are using (i.e. `ws://localhost:4000`), click "CONNECT", and make sure log says "CONNECTED". Then input a message to send. And you'll see the message sent and echoed back to us. You will also see the message logged in the terminal.
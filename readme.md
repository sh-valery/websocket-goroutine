# Web Socket Interaction

This is a simple web socket server that pushes the current time and a randomly generated UUID to the web socket every 5 seconds. And there is a web page with a client that shows the last received message.

## Notes
### Testing
The service has dependency injection to make the code testable. It follows repository pattern and the handler layer incapsulates the service layer. 
During tests the service layer is mocked and the handler layer is tested.
Also, you can change your service layer to another one and use the project as a base.

### Interaction
Service layer run a fake goroutine that pushes messages every N second to the channel. The handler layer reads from the channel at the same time handler reads from the channel and pushes messages to to the web socket.

```
service(generate message) --(chan)--> handler(read message) --(tcp)--> web socket
```

## Running the server

To run the server, you will need to have [Go](https://golang.org/) installed on your machine.

Clone the repository and navigate to the root directory:

```bash
git clone https://github.com/sh-valery/websocket-goroutine.git
cd websocket-goroutine

```


Build and run the server:

```bash
go run ./cmd/server
```


The server will start listening on port 8080.

## Connecting to the server

To connect to the web socket, you can use the webpage provided in the `web` directory.
The webpage has a script to connect to the web socket server and after this the page shows the message received from the server.

http://localhost:8080/ follow the link to open the file served by the web server in your browser.
Or Open the HTML file directly in the `web` directory in any browser. file:///path/to/websocket-goroutine/web/example.html


## Testing the server

To run the tests, navigate to the root directory and run:

```bash
go test ./...
```

## Building and running the server in a Docker container

To build and run the server in a Docker container, you will need to have [Docker](https://www.docker.com/) installed on your machine.

Build the Docker image:

```bash
docker build -t websocket-server .
```

Run the Docker container in background:
```bash
docker run -d -p 8080:8080 --name websocket-demo  websocket-server
```

OR 
Run the Docker container in foreground and remove it after stopping:
```bash
docker run -i --rm -p 8080:8080 --name websocket-demo  websocket-server

```

open page to see the result http://localhost:8080/

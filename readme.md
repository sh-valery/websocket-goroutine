# Web Socket Interaction

This is a simple web socket server that pushes the current time and a randomly generated UUID to the web socket every 3 seconds. And a web page with a client that shows the last received message.

## Running the server

To run the server, you will need to have [Go](https://golang.org/) installed on your machine.

Clone the repository and navigate to the root directory:

```bash
git clone https://github.com/sh-valery/websocket-goroutine.git
cd websocket-server

```


Build and run the server:

```bash
go run ./cmd/server
```


The server will start listening on port 8080.

## Connecting to the server

To connect to the server, you can use the webpage provided in the `web` directory.

1. Open example file served by the web server in your browser: http://localhost:8080/
or Open the HTML file directly in the `web` directory in any browser.
There is a webpage that connects to the web socket server and displays the message received from the server.

## Testing the server

To run the tests for the server, navigate to the root directory and run:

```bash
go test
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
Run the Docker container in foreground and remove it after it stops:
```bash
docker run -i --rm -p 8080:8080 --name websocket-demo  websocket-server

```

open page to see the result http://localhost:8080/

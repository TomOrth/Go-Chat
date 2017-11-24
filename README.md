# go-chat
TCP/IP chat program using golang

## Fun facts
*Utilizes custom doubly linked lists to allow for fast insertion and deletion of objects
*Utilizes `termui` to create an interactive terminal UI
*Concurrent to handle multiple requests

## Running the server
To run the server, you first have to install `termui` and the lists package used in this program by doing the following:

```
go get github.com/gizak/termui
```
```
go get github.com/TomOrth/go-chat/lists
```
Then, go to the `server` folder and run `go run main.go -port <port>`. -port is an optional flag that takes in a port number for the server to run on, but can be left out.

## Running the client
To run the client, follow the steps to install `termui` and the lists pacakge if you have not done so yet and then go to the `client` folder and run `go run main.go -port <port> -host <host> -name <name>`.  Similarly to the server, those flags are optional and can be left out.

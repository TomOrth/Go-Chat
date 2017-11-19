package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", "localhost:8000")
  go func() {
    for {
      // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("\nMessage from server: "+message) 
    }
  }()
  for { 
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text + "\n")
  }
}
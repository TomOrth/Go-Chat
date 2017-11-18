package main
import(
	"../utils"
	"fmt"
	"net"
	"bufio"
)


type Server struct {
	ServerList *utils.List
}

func (s *Server) ListUsers() {
	temp := s.ServerList.Head
	for temp != nil {
		fmt.Println(temp.Value.Name, temp.Value.Timestamp)
		temp = temp.Next
	}
}

func (s *Server) Add(user utils.User) {
	s.ServerList.Append(user)
}

var (
	connections []net.Conn
)
func serve() {
	ln, _ := net.Listen("tcp", ":6000")
    defer ln.Close()
	for {	
		// accept connection on port
		conn, _ := ln.Accept()
		connections = append(connections, conn)
		go handleRequest(conn)
	  }
}

func handleRequest(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		broadcast(message)		
	}
}

func broadcast(message string) {
	for _, client := range connections {
		client.Write([]byte(message))
	}
}
func main() {
	serve()
}


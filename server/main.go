//Package main represents the main package for the server
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"../lists"
	"github.com/gizak/termui"
)

var (
	connections *termui.Par //list of connections
)

//Type Server represents the server that clients connect to
type Server struct {
	ServerList *lists.ConnList //list of connections
}

//Add takes an incoming connection and adds it to the list
func (s *Server) Add(conn net.Conn) {
	s.ServerList.Append(conn)
}

//Delete takes a connection and removes it from the list
func (s *Server) Delete(conn net.Conn) {
	s.ServerList.Delete(conn)
}

//Serve takes a port number as a string and serves the server and handles the terminal UI
func (s *Server) Serve(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	errTer := termui.Init()
	if errTer != nil {
		panic(err)
	}
	defer termui.Close()
	connections = termui.NewPar("Connections: " + fmt.Sprint(s.ServerList.Size))
	connections.Height = 1
	connections.Width = 20
	connections.Y = 1
	connections.Border = false
	termui.Render(connections)

	for {
		// accept connection on port
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		connections.Text = "Connections: " + fmt.Sprint(s.ServerList.Size)
		termui.Render(connections)
		go s.handleRequest(conn)
		termui.Handle("/sys/kbd/q", func(termui.Event) {
			termui.StopLoop()
			os.Exit(0)
		})
	}
	termui.Loop()
}

//handleRequest takes a connection and listens on it for activity, like joining, messages, and disconnecting
func (s *Server) handleRequest(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		index := strings.Index(message, "un:")
		kindex := strings.Index(message, "/kill/")
		if index >= 0 {
			message = strings.Replace(message, "un:", "", -1)
			s.Add(conn)
			s.ServerList.Broadcast("Connected: " + message)
			connections.Text = "Connections: " + fmt.Sprint(s.ServerList.Size)
			termui.Render(connections)
		} else if kindex >= 0 {
			s.Delete(conn)
			connections.Text = "Connections: " + fmt.Sprint(s.ServerList.Size)
			s.ServerList.Broadcast("Disconnected: " + message[kindex+6:] + "\n")
			termui.Render(connections)
		} else {
			s.ServerList.Broadcast(message)
		}
	}
}

func main() {
	list := &lists.ConnList{nil, nil, 0}
	server := Server{list}
	port := flag.String("p", "8000", "port to serve on")
	flag.Parse()
	server.Serve(*port)
}

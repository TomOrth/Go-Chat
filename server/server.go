package main
import(
	"../utils"
	"net"
	"bufio"
	"flag"
	"github.com/gizak/termui"
	"fmt"
)


type Server struct {
	ServerList *utils.ConnList
}

func (s *Server) Add(user utils.User) {
	s.ServerList.Append(user)
}

//var (
//	connections []net.Conn
//)
func (s *Server) serve(port string) {
	ln, _ := net.Listen("tcp", ":" + port)
	defer ln.Close()
	
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()
	par0 := termui.NewPar("Connections: " + fmt.Sprint(s.ServerList.Size))
	par0.Height = 1
	par0.Width = 20
	par0.Y = 1
	par0.Border = false

	termui.Render(par0)

	
	termui.Handle("/timer/1s", func(e termui.Event) {
		// accept connection on port
		conn, _ := ln.Accept()
		user := utils.User{"tom", conn}		
		s.Add(user)
		par0.Text = "Connections: " + fmt.Sprint(s.ServerList.Size)
		termui.Render(par0)
		go s.handleRequest(conn)

	})

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Loop()
}

func (s *Server) handleRequest(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
	    s.ServerList.Broadcast(message)	
	}
}

func main() {
	list := &utils.ConnList{nil, nil, 0}
	server := Server{list}
	port := flag.String("p", "8000", "port to serve on")
	flag.Parse()
	server.serve(*port)
}
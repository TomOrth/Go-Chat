//Package main represents the main package for the client
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"

	"../lists"
	"github.com/gizak/termui"
)

//Type Client represents the client to connect to the server
type Client struct {
	MsgList *lists.MsgList //list of messages
	conn    net.Conn       //connection to server
	text    string         //message to be sent
	name    string         //name of client
}

//Add takes a message and adds it to the list of messages
func (c *Client) Add(msg string) {
	c.MsgList.Append(msg)
}

var (
	input *termui.Par  //input for user
	ls    *termui.List //list to show msgs
)

//Kbd handles majority of the keyboard input for the user to display a message
func (c *Client) Kbd() {
	termui.Handle("/sys/kbd", func(e termui.Event) {
		c.text += e.Data.(termui.EvtKbd).KeyStr
		input.Text = ">" + c.text
		termui.Render(termui.Body)
	})
}

//Entr handles the enter click by the user, sending the message to the server
func (c *Client) Entr() {
	termui.Handle("/sys/kbd/<enter>", func(e termui.Event) {
		fmt.Fprintf(c.conn, c.name+": "+c.text+"\n")
		c.text = ""
		input.Text = ">" + c.text
		termui.Render(termui.Body)
	})
}

//BackSp handles the backspace click by the user, deleting a character in the message
func (c *Client) BackSp() {
	termui.Handle("/sys/kbd/<backspace>", func(e termui.Event) {
		sz := len(c.text)
		if sz > 0 {
			c.text = c.text[:sz-1]
			input.Text = ">" + c.text
			termui.Render(termui.Body)
		}
	})
}

//Spce handles the space click by the user, adding a space
func (c *Client) Spce() {
	termui.Handle("/sys/kbd/<space>", func(e termui.Event) {
		c.text += " "
		input.Text = ">" + c.text
		termui.Render(termui.Body)
	})
}

//Close handles the ctrl-c click by the user, terminating the client and letting the client know
func (c *Client) Close() {
	termui.Handle("/sys/kbd/C-c", func(termui.Event) {
		c.conn.Write([]byte("/kill/" + c.name))
		termui.StopLoop()
	})
}

//Listen handles listening for incoming messages from the server
func (c *Client) Listen() {
	for {
		// listen for reply
		message, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		c.Add(strings.Replace(message, "\n", "", -1))
		if c.MsgList.Size > 20 {
			c.MsgList.DeleteHead()
		}
		ls.Items = c.MsgList.MessageArr()
		termui.Render(termui.Body)
	}
}

//Run starts up the client and creates the terminal UI, passing a username, host and port
func (c *Client) Run(name, host, port string) {
	// connect to this socket
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("un:" + name + "\n"))
	c.conn = conn
	errTer := termui.Init()
	if errTer != nil {
		panic(err)
	}
	defer termui.Close()
	input = termui.NewPar(">")
	input.Height = 1
	input.Border = false

	ls = termui.NewList()
	ls.Items = c.MsgList.MessageArr()
	ls.Overflow = "wrap"
	ls.ItemFgColor = termui.ColorYellow
	ls.Height = 22

	termui.Body.AddRows(
		termui.NewRow(termui.NewCol(12, 0, ls)),
		termui.NewRow(termui.NewCol(12, 0, input)))
	termui.Body.Align()
	termui.Render(termui.Body)

	c.Kbd()
	c.Entr()
	c.BackSp()
	c.Spce()
	c.Close()

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		termui.Body.Width = termui.TermWidth()
		termui.Body.Align()
		termui.Clear()
		termui.Render(termui.Body)
	})

	go c.Listen()
	termui.Loop()
}

func main() {
	name := flag.String("name", "tom", "username")
	host := flag.String("host", "localhost", "host name")
	port := flag.String("port", "8000", "port number")
	flag.Parse()
	messages := &lists.MsgList{nil, nil, 0}
	c := &Client{messages, nil, "", *name}
	c.Run(*name, *host, *port)
}

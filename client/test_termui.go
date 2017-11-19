package main

import "net"
import "fmt"
import "bufio"
import "github.com/gizak/termui"

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "localhost:8000")
	messages := make([]string, 0)
	text := ""
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()
	par0 := termui.NewPar(">")
	par0.Height = 1
	par0.Width = 20
	par0.Y = 9
	par0.Border = false

	ls := termui.NewList()
	ls.Items = messages
	ls.ItemFgColor = termui.ColorYellow
	ls.Height = 8
	ls.Width = 120
	ls.Y = 0

	termui.Render(par0, ls)
	termui.Handle("/sys/kbd", func(e termui.Event) {
		text += e.Data.(termui.EvtKbd).KeyStr
		par0.Text = ">" + text
		termui.Render(par0, ls)
	})

	termui.Handle("/sys/kbd/<enter>", func(e termui.Event) {
		fmt.Fprintf(conn, text+"\n")
		text = ""
		par0.Text = ">" + text
		termui.Render(par0, ls)
	})

	termui.Handle("/sys/kbd/<backspace>", func(e termui.Event) {
		sz := len(text)

		if sz > 0 {
			text = text[:sz-1]
			par0.Text = ">" + text
			termui.Render(par0, ls)
		}
	})

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	go func() {
		for {
			// listen for reply
			message, _ := bufio.NewReader(conn).ReadString('\n')
			size := len(message)
			messages = append(messages, message[:size-1])
			termui.Render(par0, ls)
			ls.Items = messages
			//termui.Render(par0)
		}
	}()

	termui.Loop()
}

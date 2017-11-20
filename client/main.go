package main

import (
	"bufio"
	"fmt"
	"net"

	"../utils"
	"github.com/gizak/termui"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "localhost:8000")
	messages := &utils.MsgList{nil, nil, 0}
	text := ""
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()
	par0 := termui.NewPar(">")
	par0.Height = 1
	par0.Border = false

	ls := termui.NewList()
	ls.Items = messages.MessageArr()
	ls.Overflow = "wrap"
	ls.ItemFgColor = termui.ColorYellow
	ls.Height = 22

	termui.Body.AddRows(
		termui.NewRow(termui.NewCol(12, 0, ls)),
		termui.NewRow(termui.NewCol(12, 0, par0)))
	termui.Body.Align()
	termui.Render(termui.Body)
	termui.Handle("/sys/kbd", func(e termui.Event) {
		text += e.Data.(termui.EvtKbd).KeyStr
		par0.Text = ">" + text
		termui.Render(termui.Body)
	})

	termui.Handle("/sys/kbd/<enter>", func(e termui.Event) {
		fmt.Fprintf(conn, text+"\n")
		text = ""
		par0.Text = ">" + text
		termui.Render(termui.Body)
	})

	termui.Handle("/sys/kbd/<backspace>", func(e termui.Event) {
		sz := len(text)

		if sz > 0 {
			text = text[:sz-1]
			par0.Text = ">" + text
			termui.Render(termui.Body)
		}
	})

	termui.Handle("/sys/kbd/<space>", func(e termui.Event) {
		text += " "
		par0.Text = ">" + text
		termui.Render(termui.Body)
	})

	termui.Handle("/sys/kbd/C-c", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		termui.Body.Width = termui.TermWidth()
		termui.Body.Align()
		termui.Clear()
		termui.Render(termui.Body)
	})

	go func() {
		for {
			// listen for reply
			message, _ := bufio.NewReader(conn).ReadString('\n')
			size := len(message)
			messages.Append(message[:size-1])
			if messages.Size > 24 {
				messages.DeleteHead()
			}
			ls.Items = messages.MessageArr()
			termui.Render(termui.Body)
		}
	}()
	termui.Loop()
}

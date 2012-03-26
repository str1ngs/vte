package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
	"vte"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	terminal := NewTerminal()
	terminal.Fork("bash")
	terminal.Connect("child-exited", gtk.MainQuit)
	window.Add(terminal)
	window.ShowAll()
	gtk.Main()
}
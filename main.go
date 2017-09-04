package main

// Sequente
// One Track Multi-Channel MIDI Sequencer
// 20170905
// M Arias


import (
	"fmt"
	"time"
	"github.com/gotk3/gotk3/gtk"
	//"github.com/rakyll/portmidi"
)


func main() {

	go Midi()

    gtk.Init(nil)

	mainWin()

	gtk.Main()

}


func mainWin() {
	mwin, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	mwin.SetPosition(gtk.WIN_POS_CENTER)
	mwin.SetDefaultSize(800, 600)
	mwin.SetTitle("Sequente 13g")
	mwin.AddEvents(2097152)

	mwin.Connect("destroy", func() {
		gtk.MainQuit()
	})

	mwin.ShowAll()
}


func Midi() {
	n := 0

	for n < 100 {
		fmt.Println("loop:", n)

		time.Sleep(200 * time.Millisecond)
		n++
	}
}


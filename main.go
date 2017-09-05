package main

// Sequente
// One Track Multi-Channel MIDI Sequencer
// 20170905
// M Arias

import (
	"fmt"
	"time"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rakyll/portmidi"
)

type Status struct {
	MidiOn bool
}


func main() {

	s := Status{true}

	go Midi(&s)

	gtk.Init(nil)

	mainWin(&s)

	gtk.Main()

}


func Midi(s *Status) {
	n := 1

	portmidi.Initialize()
	_, err := portmidi.NewOutputStream(portmidi.DefaultOutputDeviceID(), 1024, 0)
	if err != nil {
		fmt.Println("error opening output stream -", err)
		portmidi.Terminate()
		return
	}

	for s.MidiOn == true {
		fmt.Println("loop:", n)
		time.Sleep(200 * time.Millisecond)
		n++
	}

	fmt.Println("terminating portmidi")
	portmidi.Terminate()
}


func mainWin(s *Status) {
	mwin, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	mwin.SetPosition(gtk.WIN_POS_CENTER)
	mwin.SetDefaultSize(500, 300)
	mwin.SetTitle("Sequente 13g")
	mwin.AddEvents(2097152)

	mwin.Connect("destroy", func() {
		gtk.MainQuit()
	})

	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	mwin.Add(vbox)

	check, _ := gtk.SwitchNew()
	check.SetActive(s.MidiOn)
	check.SetCanFocus(false)
	check.SetMarginEnd(8)
	vbox.Add(check)

	check.Connect("state-set", func() {
		che := check.GetActive()
		if che == false {
			s.MidiOn = false
		}
	})

	mwin.ShowAll()
}


package main

// Sequente
// One Track Multi-Channel MIDI Sequencer
// 20170905
// M Arias

import (
	"fmt"
	"strconv"
	"time"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rakyll/portmidi"
)

type Status struct {
	MidiOn bool
	Count  int
	Label  *gtk.Label
}


func main() {

	s := Status{true, 0, nil}

	gtk.Init(nil)

	s.Label, _ = gtk.LabelNew("Loop: " + strconv.Itoa(s.Count))

	go Midi(&s)

	mainWin(&s)

	gtk.Main()

}


func Midi(s *Status) {
	
	portmidi.Initialize()
	_, err := portmidi.NewOutputStream(portmidi.DefaultOutputDeviceID(), 1024, 0)
	if err != nil {
		fmt.Println("error opening output stream -", err)
		portmidi.Terminate()
		return
	}

	for s.MidiOn == true {
		s.Count++
		s.Label.SetText("Loop: " + strconv.Itoa(s.Count))
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("terminating portmidi")
	portmidi.Terminate()
}


func mainWin(s *Status) {
	mwin, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	mwin.SetPosition(gtk.WIN_POS_CENTER)
	mwin.SetDefaultSize(300, 150)
	mwin.SetBorderWidth(8)
	mwin.SetTitle("Sequente 13g")
	mwin.AddEvents(2097152)

	mwin.Connect("destroy", func() {
		gtk.MainQuit()
	})

	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	mwin.Add(vbox)
	hbox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	vbox.Add(hbox)	
	hbox.PackStart(s.Label, false, false, 8)

	check, _ := gtk.SwitchNew()
	check.SetActive(s.MidiOn)
	check.SetCanFocus(false)
	check.SetMarginEnd(8)
	hbox.PackEnd(check, false, false, 0)

	check.Connect("state-set", func() {
		che := check.GetActive()
		if che == false {
			s.MidiOn = false
		}
	})

	mwin.ShowAll()
}


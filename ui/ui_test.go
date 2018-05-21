package ui

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
	"testing"
)

func TestRun(t *testing.T) {
	New(nil, 3, 3)
	go Run()
	termbox.Interrupt()
}

func TestUIBorder(t *testing.T) {
	New(nil, 3, 3)
	SetBorder(LightBorder, true)
	Init()

	runes, err := PrintScreen()
	if err != nil {
		t.Fatal(err)
	}
	termbox.Close()

	want := "┌─┐\n│ │\n└─┘\n"
	if string(runes) != want {
		fmt.Printf("Wanted:\"\n%s\"\nGot:\"\n%s\"\n", want, string(runes))
		t.Fatal(t)
	}
}

func TestBorderSet_Invisible(t *testing.T) {
	New(nil, 3, 3)
	SetBorder(LightBorder, false)
	Init()

	runes, err := PrintScreen()
	if err != nil {
		t.Fatal(err)
	}
	termbox.Close()

	want := "   \n   \n   \n"
	if string(runes) != want {
		t.Fatalf("Wanted:\"\n%s\"\nGot:\"\n%s\"\n", want, string(runes))
	}
}

func TestPrint(t *testing.T) {
	want := "Two   \n!ines!\n"

	New(nil, 6, 2)
	Init()
	Print(0, 0, "Two\nLines!\r!")
	runes, err := PrintScreen()
	if err != nil {
		t.Fatal(err)
	}
	termbox.Close()

	if string(runes) != want {
		t.Fatalf("Wanted:\"\n%s\"\nGot:\"\n%s\"\n", want, string(runes))
	}
}

func TestPrintRaw(t *testing.T) {
	want := "Two Lines! !\n"

	New(nil, 12, 1)
	Init()
	PrintRaw(0, 0, "Two\nLines!\r!")
	runes, err := PrintScreen()
	if err != nil {
		t.Fatal(err)
	}
	termbox.Close()

	if string(runes) != want {
		t.Fatalf("Wanted:\"\n%s\"\nGot:\"\n%s\"\n", want, string(runes))
	}
}

func TestFullScreen(t *testing.T) {
	want := `┌────────────────────────────────────────────────────────────────────────────────┐
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
│................................................................................│
└────────────────────────────────────────────────────────────────────────────────┘
`
	Standard(nil, engine.NewMap(80, 24, "Map"))
	Init()
	runes, err := PrintScreen()
	if err != nil {
		t.Fatal(err)
	}
	termbox.Close()

	if string(runes) != want {
		t.Fatalf("Wanted:\"\n%s\"\nGot:\"\n%s\"\n", want, string(runes))
	}
}

/*
func TestTextField(t *testing.T) {
	Standard(nil, engine.NewMap(80, 24, "Map"))
	Init()
	w, h := Size()
	v := NewTextField("cmds", NewBounds(0, h-3, w, h-1))
	_ = termbox.PollEvent()
	v.Popup()
	termbox.Close()
	fmt.Println(v.Text())
}
*/

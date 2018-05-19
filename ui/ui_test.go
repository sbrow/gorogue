package ui

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	"strings"
	"testing"
)

func TestBorderSet_Light(t *testing.T) {
	if err := termbox.Init(); err != nil {
		t.Fatal(err)
	}
	defer termbox.Close()

	// Top-Left corner.
	x1, y1 := 0, 0
	// Bottom-Right corner.
	x2, y2 := 3, 3

	b := NewBorder(LightBorder, true)
	bnds := NewBounds(x1, y1, x2, y2)
	b.Draw(bnds)
	// r := PrintScreen()
	runes, err := PrintScreen()
	if err != nil {
		t.Fatal(err)
	}
	v := CellsToByte(runes, x1, y1, x2, y2)
	fmt.Printf("v:\"\n%s\"\n", v)
	comp := "┌─┐\n│ │\n└─┘\n"
	if string(v) != comp {
		t.Fatal(fmt.Sprintf("Expected:\"\n%s\"\nGot:\"\n%s\"", comp, v))
	}
}

func TestBorderSet_Invisible(t *testing.T) {
	if err := termbox.Init(); err != nil {
		t.Fatal(err)
	}
	defer termbox.Close()

	// Top-Left corner.
	x1, y1 := 0, 0
	// Bottom-Right corner.
	x2, y2 := 3, 3

	b := NewBorder(LightBorder, false)
	bnds := NewBounds(x1, y1, x2, y2)
	b.Draw(bnds)
	// r := PrintScreen()
	runes, err := PrintScreen()
	if err != nil {
		t.Fatal(err)
	}
	v := CellsToByte(runes, x1, y1, x2, y2)
	fmt.Printf("v:\"\n%s\"\n", v)
	comp := "   \n   \n   \n"
	if string(v) != comp {
		t.Fatal(fmt.Sprintf("Expected:\"\n%s\"\nGot:\"\n%s\"", comp, v))
	}
}

func TestCells(t *testing.T) {
	if err := termbox.Init(); err != nil {
		t.Fatal(err)
	}
	defer termbox.Close()
	w, h := termbox.Size()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			termbox.SetCell(x, y, '.', termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	runes, err := PrintScreen()
	if err != nil {
		t.Fatal(err)
	}
	v := CellsToByte(runes, 0, 0, 4, 4)
	comp := []rune{}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			comp = append(comp, '.')
		}
		comp = append(comp, '\n')
	}

	if string(v) != string(comp) {
		t.Fatal(fmt.Sprintf("Expected: %s\nGot: %s\n", string(comp), v))
	}
}

func TestDrawAt(t *testing.T) {
	if err := termbox.Init(); err != nil {
		t.Fatal(err)
	}
	defer termbox.Close()

	str := "Two\nLines!\r!"
	// TODO: Incomplete test
	tests := []interface{}{
		str,
		strings.Split(str, "\n"),
		[]rune(str),
		[]byte(str),
	}
	x1, y1 := 0, 0
	x2, y2 := 6, 2

	for _, test := range tests {
		DrawAt(0, 0, test)

		runes, err := PrintScreen()
		if err != nil {
			t.Fatal(err)
		}
		v := CellsToByte(runes, x1, y1, x2, y2)
		out := "Two   \n!ines!\n"
		if string(v) != out {
			t.Fatalf("Expected:\"\n%s\"\nGot:\"\n%s\"", out, v)
		}
	}
}

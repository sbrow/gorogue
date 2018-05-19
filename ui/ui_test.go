package ui

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
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
	v := things(PrintScreen(), x1, y1, x2, y2)
	fmt.Printf("v:\"\n%s\"\n", v)
	comp := "┌─┐\n│ │\n└─┘\n"
	if v != comp {
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
	v := things(PrintScreen(), 0, 0, 4, 4)
	comp := []rune{}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			comp = append(comp, '.')
		}
		comp = append(comp, '\n')
	}

	if v != string(comp) {
		t.Fatal(fmt.Sprintf("Expected: %s\nGot: %s\n", string(comp), v))
	}
}

// func PrintScreen() ([][]rune, error) {
func PrintScreen() [][]rune {
	w, h := termbox.Size()
	fmt.Println("w", w, "h", h)
	if w == 0 || h == 0 {
		panic("Termbox has no size. Has it been initialized?")
	}
	cells := termbox.CellBuffer()
	runes := [][]rune{}
	for x := 0; x < w; x++ {
		runes = append(runes, []rune{})
		for y := 0; y < h; y++ {
			runes[x] = append(runes[x], cells[(y*w)+x].Ch)
		}
	}
	return runes
}

func things(runes [][]rune, x1, y1, x2, y2 int) string {
	str := ""
	for y := y1; y < y2; y++ {
		for x := x1; x < x2; x++ {
			str += string(runes[x][y])
		}
		str += "\n"
	}
	return str
}

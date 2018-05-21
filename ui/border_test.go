package ui

import (
	termbox "github.com/nsf/termbox-go"
	// "testing"
)

func DrawBorderSet(b BorderSet, bounds Bounds) {
	defer termbox.Flush()

	// Top-Left corner.
	x1, y1 := bounds[0].Ints()
	// Bottom-Right corner.
	x2, y2 := bounds[1].Ints()

	s := []rune(string(b))

	// Print the horizontals
	for x := x1; x < x2; x++ {
		termbox.SetCell(x, y1, s[Horizontal], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x, y2/2, s[Horizontal], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x, y2, s[Horizontal], termbox.ColorDefault, termbox.ColorBlack)
	}
	// Print the verticals
	for y := y1; y < y2; y++ {
		termbox.SetCell(x1, y, s[Vertical], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x2/2, y, s[Vertical], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x2, y, s[Vertical], termbox.ColorDefault, termbox.ColorBlack)
	}

	// Print the corners.
	termbox.SetCell(x1, y1, s[LeftTop], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2/2, y1, s[MiddleTop], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2, y1, s[RightTop], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x1, y2/2, s[LeftMiddle], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2/2, y2/2, s[Center], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2, y2/2, s[RightMiddle], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x1, y2, s[LeftBottom], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2/2, y2, s[MiddleBottom], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2, y2, s[RightBottom], termbox.ColorDefault, termbox.ColorBlack)
}

/*func TestBorderSet_Light(t *testing.T) {
	New(nil, 5, 5)
	Init()
	DrawBorderSet(LightBorder, *OuterBounds())
	// _ = termbox.PollEvent()

	runes, err := PrintScreen()
	if err != nil {
		t.Fatal(err)
	}
	termbox.Close()

	want := "┌─┐\n│ │\n└─┘\n"
	if string(runes) != want {
		t.Fatal("butts")
		// t.Fatalf("Wanted:\"\n%s\"\nGot:\"\n%s\"\n", want, string(runes))
	}
}*/

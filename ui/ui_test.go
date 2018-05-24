package ui_test

import (
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/assets"
	"github.com/sbrow/gorogue/ui"
	"testing"
)

func Frame() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	ui.Draw()
}

func TestRun(t *testing.T) {
	ui.Init(3, 3)
	go ui.Run()
	termbox.Interrupt()
}

func TestUIBorder(t *testing.T) {
	ui.Init(3, 3)
	ui.SetBorder(ui.LightBorder, true)
	Frame()

	runes, err := ui.PrintScreen()
	if err != nil {
		t.Fatal(err)
	}
	termbox.Close()

	want := "┌─┐\n│ │\n└─┘\n"
	if string(runes) != want {
		t.Fatalf("Wanted:\"\n%s\"\nGot:\"\n%s\"\n", want, string(runes))
	}
}

func TestBorderSet_Invisible(t *testing.T) {
	ui.Init(3, 3)
	ui.SetBorder(ui.LightBorder, false)
	Frame()

	runes, err := ui.PrintScreen()
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

	ui.Init(6, 2)
	Frame()
	ui.Print(0, 0, "Two\nLines!\r!")
	runes, err := ui.PrintScreen()
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

	ui.Init(12, 1)
	Frame()
	ui.PrintRaw(0, 0, "Two\nLines!\r!")
	runes, err := ui.PrintScreen()
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
	ui.Init(82, 26)
	tiles := engine.NewMap(80, 24).AllTiles()
	assets.StandardUI(&tiles)
	Frame()
	// Frame()
	runes, err := ui.PrintScreen()
	if err != nil {
		t.Fatal(err)
	}
	termbox.Close()

	if string(runes) != want {
		t.Fatalf("Wanted:\"\n%s\"\nGot:\"\n%s\"\n", want, string(runes))
	}
}

// TODO: Doesn't compare to anything.
/*func TestTextField(t *testing.T) {
	engine.SetLog("testing.log", true)
	// engine.FrameClient(&example.Client{})
	tiles := engine.FrameMap(80, 24).AllTiles()
	assets.StandardUI(&tiles)
	Frame()
	w, h := Size()
	v := FrameTextField("cmds", FrameBounds(0, h-3, w, h-1))
	_ = termbox.PollEvent()
	v.Popup()
	termbox.Close()
	fmt.ui.Println(v.Text())
}
*/

func DrawBorderSet(b ui.BorderSet, bounds ui.Bounds) {
	defer termbox.Flush()

	// Top-Left corner.
	x1, y1 := bounds[0].Ints()
	// Bottom-Right corner.
	x2, y2 := bounds[1].Ints()

	s := []rune(string(b))

	// ui.Print the horizontals
	for x := x1; x < x2; x++ {
		termbox.SetCell(x, y1, s[ui.Horizontal], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x, y2/2, s[ui.Horizontal], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x, y2, s[ui.Horizontal], termbox.ColorDefault, termbox.ColorBlack)
	}
	// ui.Print the verticals
	for y := y1; y < y2; y++ {
		termbox.SetCell(x1, y, s[ui.Vertical], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x2/2, y, s[ui.Vertical], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x2, y, s[ui.Vertical], termbox.ColorDefault, termbox.ColorBlack)
	}

	// ui.Print the corners.
	termbox.SetCell(x1, y1, s[ui.LeftTop], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2/2, y1, s[ui.MiddleTop], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2, y1, s[ui.RightTop], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x1, y2/2, s[ui.LeftMiddle], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2/2, y2/2, s[ui.Center], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2, y2/2, s[ui.RightMiddle], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x1, y2, s[ui.LeftBottom], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2/2, y2, s[ui.MiddleBottom], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2, y2, s[ui.RightBottom], termbox.ColorDefault, termbox.ColorBlack)
}

func TestBorderSet_Light(t *testing.T) {
	if err := termbox.Init(); err != nil {
		t.Fatal(err)
	}
	ui.Init(5, 5)
	ui.Draw()
	DrawBorderSet(ui.LightBorder, *ui.OuterBounds())
	// _ = termbox.PollEvent()

	runes, err := ui.PrintScreen()
	if err != nil {
		t.Fatal(err)
	}
	termbox.Close()

	want := `┌─┬─┐
│ │ │
├─┼─┤
│ │ │
└─┴─┘
`
	if string(runes) != want {
		t.Fatalf("Wanted:\"\n%s\"\nGot:\"\n%s\"\n", want, string(runes))
	}
}

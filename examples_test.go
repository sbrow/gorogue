package gorogue

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	"testing"
)

func ExampleActionQuit_Action() {
	quit := ActionQuit{"Client"}
	action := quit.Action()

	fmt.Println(action.Name(), action.Caller(), action.Args())
	// Output: Quit Client []
}

func TestDrawString(t *testing.T) {
	termbox.Init()
	defer termbox.Close()
	x, y := 5, 10
	str := "This string\nreturns\rbreak."
	DrawString(x, y, termbox.ColorDefault, termbox.ColorDefault, str)
	_ = termbox.PollEvent()
}

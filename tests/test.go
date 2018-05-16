package main

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	"github.com/sbrow/gorogue/keys"
	. "github.com/sbrow/gorogue/ui"
)

func main() {
	termbox.Init()
	w, h := termbox.Size()
	// defer termbox.Close()
	// v := New("cmds", NewBounds(1, h-1, w-1, h-1))
	v := New("cmds", NewBounds(1, h-2, w-1, h-1))
	v.Draw()
	termbox.Close()
	fmt.Println(v.Text())
}

type TextField struct {
	name   string
	border BorderSet
	bounds Bounds
	text   string
}

func New(name string, b Bounds) *TextField {
	return &TextField{
		name:   name,
		border: HeavyBorder,
		bounds: b,
		text:   ":",
	}
}

func (t *TextField) Bounds() Bounds {
	return t.bounds
}

func (t *TextField) Draw() {
	defer termbox.Flush()
	t.border.Draw(t.bounds)
	x, y := t.bounds[0].Ints()
	zero := 2
	for _, r := range t.text {
		termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
	termbox.SetCursor(x, y)
main:
	for {
		termbox.SetCursor(x, y)
		termbox.Flush()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			k := &keys.Key{}
			k.Mod = ev.Mod
			if ev.Ch != 0 {
				k.Ch = ev.Ch
			} else {
				k.Key = ev.Key
			}
			switch {
			case *k == keys.Esc:
				t.text = ":"
				break main
			case *k == keys.Enter:
				break main
			case *k == keys.Backspace:
				fallthrough
			case k.Key == termbox.KeyBackspace2:
				fallthrough
			case *k == keys.Delete:
				if x > zero {
					x--
					termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
				}
			case *k == keys.Space:
				k.Ch = ' '
				fallthrough
			case k.Mod == 0 && k.Key == 0:
				termbox.SetCell(x, y, k.Ch, termbox.ColorDefault, termbox.ColorDefault)
				x++
			}
		}
	}
	x1, _ := t.bounds[0].X
	x2, _ := t.bounds[1].X
	t.text = ""
	for x := x1; x <= x2; x++ {

	}
}

func (t *TextField) Name() string {
	return t.name
}

func (t TextField) Text() string {
	return t.text
}
func (t TextField) Type() UIElementType {
	return UITypeTextField
}

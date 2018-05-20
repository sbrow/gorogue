package ui

import (
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
)

type TextField struct {
	name   string
	border *Border
	bounds Bounds
	text   string
	prefix string
}

func NewTextField(name string, b Bounds) *TextField {
	return &TextField{
		name:   name,
		border: NewBorder(HeavyBorder, true),
		bounds: b,
		prefix: ":",
	}
}

func (t *TextField) Bounds() Bounds {
	return t.bounds
}

func (t *TextField) Draw() {
	defer termbox.Flush()
	t.border.Draw(t.bounds)
	x, y := t.bounds[0].Ints()
	if t.border.Visible {
		x++
		y++
	}
	for _, r := range t.prefix + t.text {
		termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
	termbox.SetCursor(x, y)
	for x < t.bounds[1].X {
		termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
}

func (t *TextField) Popup() {
	defer termbox.Flush()
	x, y := t.bounds[0].Ints()
	if t.border.Visible {
		x++
		y++
	}
	zero := x
	x += len(t.text)
main:
	for {
		Draw()
		t.Draw()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			k := &engine.Key{}
			k.Mod = ev.Mod
			if ev.Ch != 0 {
				k.Ch = ev.Ch
			} else {
				k.Key = ev.Key
			}
			switch {
			case *k == engine.Esc:
				t.text = ""
				break main
			case *k == engine.Enter:
				break main
			case *k == engine.Backspace:
				fallthrough
			case k.Key == termbox.KeyBackspace2:
				fallthrough
			case *k == engine.Delete:
				if x > zero {
					x--
					termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
					t.text = t.text[:len(t.text)-1]
				}
			case *k == engine.Space:
				k.Ch = ' '
				fallthrough
			case k.Mod == 0 && k.Key == 0:
				termbox.SetCell(x, y, k.Ch, termbox.ColorDefault, termbox.ColorDefault)
				t.text += string(k.Ch)
				x++
			}
		}
	}
	t.text = t.prefix + t.text
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

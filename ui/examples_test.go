package ui

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
)

func ExamplePrint() {
	// Set up the UI
	termbox.Init()
	defer termbox.Close()
	Init(16, 4)
	SetBorder(HeavyBorder, true)
	Draw()

	str := "Hello,\nWorld\r!"
	Print(1, 1, str)

	data, err := PrintScreen()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	// Output:
	// ┏━━━━━━━━━━━━━━┓
	// ┃Hello,        ┃
	// ┃!orld         ┃
	// ┗━━━━━━━━━━━━━━┛
}

func ExamplePrintRaw() {
	// Set up the UI
	termbox.Init()
	defer termbox.Close()
	Init(16, 4)
	SetBorder(HeavyBorder, true)
	Draw()

	str := "Hello,\nWorld\r!"
	PrintRaw(1, 1, str)

	data, err := PrintScreen()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// Output:
	// ┏━━━━━━━━━━━━━━┓
	// ┃Hello, World !┃
	// ┃              ┃
	// ┗━━━━━━━━━━━━━━┛
}

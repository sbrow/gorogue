package gorogue

import (
	"fmt"
)

func ExampleActionQuit_Action() {
	quit := ActionQuit{"Client"}
	action := quit.Action()

	fmt.Println(action.Name(), action.Caller(), action.Args())
	// Output: Quit Client []
}

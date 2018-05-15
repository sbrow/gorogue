package assets

import (
	"fmt"
	. "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/ui"
)

// UITemplate provides some example UIs to help with starting your game.
// type UITemplate struct {
// *UI
// }

// Fullscreen returns a minimum terminal sized UI (80x24),
// with one view into a map.
func Fullscreen(c Client, m *string) UI {
	fmt.Println("Fullscreen", m)
	// Initialize back-end

	// Initialize front-end
	u := ui.NewUI(c, "Fullscreen Game", 80, 24)

	// Point to start pulling map data from.
	mapOrigin := Point{0, 0}

	// Size of our UI minus the border.
	viewSize := Point{80, 24}

	// where to place the view in the UI.
	viewOrign := Point{0, 0}

	// Create a new view into exampleMap
	v := ui.NewView(ui.Bounds{mapOrigin, viewSize}, m, viewOrign)

	// Fill our UI with a view of our map.
	u.AddView("World", *v)

	// Add a border
	u.SetBorder(ui.LightBorder)

	return u
}

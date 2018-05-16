package assets

import (
	. "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/ui"
)

// Fullscreen returns a minimum terminal sized UI (80x24),
// with one view into a map.
func Fullscreen(c Client, m Map) UI {

	// Point to start pulling map data from.
	mapOrigin := Point{0, 0}

	// Size of our UI minus the border.
	viewSize := Point{81, 25}

	// where to place the view in the UI.
	viewOrign := Point{0, 0}

	// Initialize UI
	u := ui.NewUI(c, "Fullscreen Game", viewSize.X, viewSize.Y)

	// Create a new view into exampleMap
	v := ui.NewView(ui.Bounds{mapOrigin, viewSize}, m, viewOrign)

	// Fill our UI with a view of our map.
	u.Add("Map", v)

	// Add a border
	u.SetBorder(&ui.Border{ui.LightBorder, true})

	return u
}

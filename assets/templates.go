package assets

import (
	. "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/ui"
)

// StandardUI initializes a minimum terminal sized UI (80x24),
// with one view into a map.
func StandardUI(m *[][]Tile) {

	// Point to start pulling map data from.
	origin := Point{0, 0}

	// Size of our UI.
	uiSize := Point{82, 26}

	// Initialize UI
	ui.Init(uiSize.X, uiSize.Y)

	// Add a border.
	ui.SetBorder(ui.LightBorder, true)

	// Add a view of the map to our UI.
	ui.Add("Map", ui.NewView(origin, m))
}

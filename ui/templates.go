package ui

import engine "github.com/sbrow/gorogue"

// Standard returns a minimum terminal sized UI (80x24),
// with one view into a map.
func Standard(m *[][]engine.Tile) {

	// Point to start pulling map data from.
	origin := engine.Point{0, 0}

	// Size of our UI.
	uiSize := engine.Point{82, 26}

	// Initialize UI
	New(uiSize.X, uiSize.Y)
	SetBorder(LightBorder, true)

	// Add a view of the map to our UI.
	Add("Map", NewView(origin, m))
}

package ui

import engine "github.com/sbrow/gorogue"

// Standard returns a minimum terminal sized UI (80x24),
// with one view into a map.
func Standard(c engine.Client, m *engine.Map) {

	// Point to start pulling map data from.
	mapOrigin := engine.Point{0, 0}

	// Size of our UI minus the border.
	uiSize := engine.Point{82, 26}

	viewBounds := Bounds{mapOrigin, engine.Point{80, 24}}

	// where to place the view in the UI.
	viewOrign := engine.Point{0, 0}

	// Initialize UI
	New(c, uiSize.X, uiSize.Y)

	// Create a new view into exampleMap
	v := NewView(viewBounds, m, viewOrign)

	// Fill our UI with a view of our map.
	Add("Map", v)

	// Add a border
	SetBorder(LightBorder, true)
}

package ui

import engine "github.com/sbrow/gorogue"

// Fullscreen returns a minimum terminal sized UI (80x24),
// with one view into a map.
func Fullscreen(c engine.Client, m *engine.Map) engine.UI {

	// Point to start pulling map data from.
	mapOrigin := engine.Point{0, 0}

	// Size of our UI minus the border.
	viewSize := engine.Point{81, 25}

	// where to place the view in the UI.
	viewOrign := engine.Point{0, 0}

	// Initialize UI
	u := NewUI(c, "Fullscreen Game", viewSize.X, viewSize.Y)

	// Create a new view into exampleMap
	v := NewView(Bounds{mapOrigin, viewSize}, m, viewOrign)

	// Fill our UI with a view of our map.
	u.Add("Map", v)

	// Add a border
	u.SetBorder(&Border{LightBorder, true})

	return u
}

package ui

import engine "github.com/sbrow/gorogue"

// Fullscreen returns a minimum terminal sized UI (80x24),
// with one view into a map.
func Fullscreen(c engine.Client, m *engine.Map) engine.UI {

	// Point to start pulling map data from.
	mapOrigin := engine.Point{0, 0}

	// Size of our UI minus the border.
	uiSize := engine.Point{82, 26}

	viewBounds := Bounds{mapOrigin, engine.Point{uiSize.X - 1, uiSize.Y - 1}}

	// where to place the view in the UI.
	viewOrign := engine.Point{0, 0}

	// Initialize UI
	u := NewUI(c, "Fullscreen Game", 0, 0, uiSize.X, uiSize.Y)

	// Create a new view into exampleMap
	v := NewView(viewBounds, m, viewOrign)

	// Fill our UI with a view of our map.
	u.Add("Map", v)

	// Add a border
	u.SetBorder(&Border{LightBorder, true})

	return u
}

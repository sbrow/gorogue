package gorogue

import "fmt"

func ExampleDirection_Point() {
	fmt.Println(North.Point())
	fmt.Println(SouthWest.Point())
	// Output:
	// {0 -1}
	// {-1 1}
}

func ExampleMap_TileSlice() {
	m := NewMap(3, 3)
	str := "┌─────┐\n"
	for _, row := range m.TileSlice(-1, -1, 3, 3) {
		str += "│"
		for _, cell := range row {
			str += string(cell.Sprite.Ch)
		}
		str += "│\n"
	}
	fmt.Println(str + "└─────┘\n")

	// Output:
	// ┌─────┐
	// │     │
	// │ ... │
	// │ ... │
	// │ ... │
	// │     │
	// └─────┘
}

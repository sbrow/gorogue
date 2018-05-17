// Package gorogue is a flexible roguelike engine written in Go.
// Gorogue aims to be small, versatile, and modular.
//
// Game Modes
//
// One of this project's goals is to support a wide variety of game modes. However,
// emphasis is first placed on the stability and thorough documentation of the existing
// modes.
//
// Currently, there are two game modes: online and local. Both support one Action per
// Actor per tick.
//
// Planned modes include:
// 	- "Realtime": 30 ticks per second, Actors that don't act in time are skipped.
// 	- "Squad Based": Allows each player to control more than one Actor."
// 	- "Action Points": Characters spend AP to perform actions, AP refreshes each tick.
//
// Online Mode
//
// In online mode, RemoteClients connect to a Server.
package gorogue

import (
	termbox "github.com/nsf/termbox-go"
	"net/rpc"
)

// Actor is an object that can move. There are two main kinds of actors:
// player characters and non-player characters (NPCs). The important
// distinction being that NPCs are controlled by the World, and player
// characters are controlled by Clients.
type Actor interface {
	Object             // All Actors are Objects.
	Move(pos Pos) bool // Moves the Actor to the given position.
}

// Client represents a connection to a local World.
type Client interface {
	// Addr returns a string representation of the Client's host/port.
	// Defaults to "[::1]" for local Clients.
	Addr() string

	// Init sets up the Client, generally handling spawning the player(s)
	// and setting up the UI.
	Init() error

	// HandleAction sends an Action received received from Input(), Command strings, etc
	// and sends it to the World to be evaluated.
	//
	// Failed/Illegal Actions will return an error, and generally require that another
	// Action be selected instead.
	HandleAction(a *Action) error

	// Maps pulls all the Maps from the World and returns them.
	Maps() []Map

	// Player returns the Actor that this Client is in control of.
	Player() Actor

	// Run starts the UI.
	Run()

	// TODO: Document.
	SetWorld(w World)
}

// Conn is the server-side representation of a connection to a client.
/*type Conn struct {
	Host   string // Connection data.
	Player Actor  // Actors this connection has control over.
}
*/

// Direction represents the cardinal and ordinal directions.
// North points towards the top of the screen, east points to the right, etc.
//
// Converting between coordinates and Directions is often done with Bitwise operations,
// hence why they are not laid out in perfect sequence.
type Direction uint8

const (
	North     Direction = 1 + iota // 0001
	East                           // 0010
	NorthEast                      // 0011
	West                           // 0100
	NorthWest                      // 0101
	South     Direction = 8        // 1000
	SouthEast Direction = 10       // 1010
	SouthWest Direction = 12       // 1100
)

// FIXME: Map is currently under development.
type Map interface {
	// Height() int
	// Players() []Actor
	// Tick()r
	// Tiles() [][] Tile

	TileSlice(x1, x2, w, h int) [][]Tile
	// Width() int
}

type Object interface {
	Name() string
	ID() string
	Index() int
	MarshalJSON() ([]byte, error)
	Pos() *Pos
	SetIndex(i int)
	SetPos(p *Pos)
	Sprite() Sprite
	UnmarshalJSON(data []byte) error
}

// Point represents a coordinate pair.
//
// Points are most commonly used to locate Tiles on a Map and Cells in termbox.
type Point struct {
	X int
	Y int
}

// Ints returns the point as a pair of ints.
func (p *Point) Ints() (x, y int) {
	return p.X, p.Y
}

// Pos represents the position of an Object in the World. Point holds their location
// in the Map, and Map holds the index of theirr Map in World.Maps.
type Pos struct {
	Point
	Map int
}

func NewPos(x, y, Map int) *Pos {
	return &Pos{Point{x, y}, Map}
}

// Ints returns the position as an ordered triple.
func (p *Pos) Ints() (x, y, z int) {
	return p.X, p.Y, p.Map
}

// TODO: Document
type RemoteClient interface {
	Client
	Connect(host, port string)
	Disconnect()
	SetAddr(addr string)
	SetRPC(*rpc.Client)
}

// Server is a world that accepts RemoteClient connections over TCP/IP.
//
// Actions are sent from RemoteClients to the Server via JSON RPC.
// This means that every Server method you wish RemoteClients to have access to
// must follow these criteria:
//	- the method's type is exported.
//	- the method is exported.
//	- the method has two arguments, both exported (or builtin) types.
// 	- the method's second argument is a pointer.
//	- the method has return type error.
// See the net/rpc package for more details.
type Server interface {
	World
	Conns() []string
	HandleRequests()
	// Ping(addr *string, reply *Pong)
	Port() string
	SetPort(port string)
}

type Sprite termbox.Cell

var (
	DefaultPlayer Sprite = Sprite(termbox.Cell{'@', termbox.ColorWhite, termbox.ColorBlack})
)

type Tile struct {
	Sprite
	Objects []Object
}

func NewTile(sprite termbox.Cell) Tile {
	return Tile{
		Sprite: Sprite(sprite),
	}
}

var (
	EmptyTile Tile = NewTile(termbox.Cell{' ', termbox.ColorWhite, termbox.ColorBlack})
	FloorTile      = NewTile(termbox.Cell{'.', termbox.ColorWhite, termbox.ColorBlack})
)

func (t Tile) Cell() termbox.Cell {
	return termbox.Cell(t.Sprite)
}

type UI interface {
	Run()
}

type World interface {
	Maps() []Map
	Players() map[string]Actor
}

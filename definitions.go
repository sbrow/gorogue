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
// 	- "Realtime"	 : 30 ticks per second, Actors that don't act in time are skipped.
// 	- "Squad Based"	 : Allows each player to control more than one Actor."
// 	- "Action Points": Characters spend AP to perform actions, AP refreshes each tick.
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
	Object // All Actors are Objects.
	Done()
	Move(pos *Pos) error // Moves the Actor to the given position.
	Ready()
	SetMap(m *Map)
}

// Client represents a connection to a local World.
type Client interface {
	// Addr returns a string representation of the Client's host/port.
	// Defaults to "[::1]" for local Clients.
	Addr() string

	// Init sets up the Client, generally handling spawning the player(s)
	// and setting up the UI.
	Init() error

	// HandleAction sends an Action received from: Input(), Command strings, etc.
	// and sends it to the World to be evaluated.
	//
	// Failed/Illegal Actions should return an error, and require that another
	// Action be selected instead.
	HandleAction(a *Action) error

	// Maps pulls all the Maps from the World and returns them.
	// Maps() []Map

	// Player returns the Actor that this Client is in control of.
	Player() Actor
}

// Conn is the server-side representation of a connection to a client.
/*type Conn struct {
	Host   string // Connection data.
	Player Actor  // Actors this connection has control over.
}
*/

type Object interface {
	Name() string
	ID() string
	Index() int
	// MarshalJSON() ([]byte, error)
	Map() *Map
	Pos() *Pos
	SetIndex(i int)
	SetPt(pt *Point)
	// SetPos(p *Pos)
	Sprite() termbox.Cell
	// UnmarshalJSON(data []byte) error
}

// TODO: Document
type RemoteClient interface {
	Client
	// Connect(host, port string)
	Disconnect()
	Run()
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

// Tile is square on the Map. It is the smallest increment a Map can be broken
// down into.
// Tiles generally represent immovable parts of the environment:
// walls, floors, doors, etc.
type Tile struct {
	Sprite  termbox.Cell
	Objects []Object
}

// NewTile returns a Tile with the given sprite.
func NewTile(sprite termbox.Cell) Tile {
	return Tile{
		Sprite: sprite,
	}
}

var (
	BlankTile Tile = NewTile(termbox.Cell{' ', termbox.ColorWhite, termbox.ColorBlack})
	FloorTile Tile = NewTile(termbox.Cell{'.', termbox.ColorWhite, termbox.ColorBlack})
	WallTile  Tile = NewTile(termbox.Cell{'#', termbox.ColorWhite, termbox.ColorBlack})
)

type TileSet string

const (
	Tiles TileSet = " .@#><"
)

type UI interface {
	Run()
	Draw() error
}

type World interface {
	HandleAction(a *Action, reply *string) error
	Maps() []*Map
	Players() map[string]Actor
}

// Package gorogue is a flexible roguelike engine written in Go.
// Gorogue aims to be small, versatile, and modular.
package gorogue

import (
	termbox "github.com/nsf/termbox-go"
	"net"
	"net/rpc"
)

var stdConn Client

var stdUI UI

type Action interface {
	Args() []interface{}
	Caller() string
	Name() string
}

// Actor is an object that can Move. There are two main kinds of actors:
// player characters and non-player characters (NPCs). The important
// distinction being that NPCs are controlled by the server, and Player
// characters are controlled by clients.
//
// Each NPC gets their own goroutine, meaning each acts on their own thread,
// separate from other actors. The server receives requests to act from each NPC,
// and determines whether that action is valid. If if isn't, the action is rejected
// and the Actor must choose a different action to perform. If the action is valid,
// it gets stored in in a buffer and is called during the next Map.Tick().
//
// TODO: This description is only valid for the client-server version.
type Actor interface {
	Object             // The Object interface.
	Move(pos Pos) bool // Moves the Actor to the given position.
}

type Client interface {
	Addr() string
	Disconnect()
	Init() UI
	Ping()
	HandleAction(a Action) error
	Maps() map[string]Map
	Squad() []Actor
	SetAddr(addr string)
	SetRPC(conn *rpc.Client)
}

// Conn is the server-side representation of a connection to a client.
type Conn struct {
	Conn  *net.Conn // Connection data.
	Squad []Actor   // Actors this connection has control over.
}

// Direction represents the cardinal and ordinal directions.
// North points towards the top of the screen, east points to the right, etc.
//
// Converting between coordinates and Directions is often done with Bitwise operations,
// hence why they are not laid out in perfect sequence.r
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

type Map interface {
	TileSlice(x1, x2, w, h int) [][]Tile
}

type Object interface {
	Name() string
	ID() string
	Index() int
	MarshalJSON() ([]byte, error)
	Pos() *Pos
	SetIndex(i int)
	SetPos(p Pos)
	Sprite() Sprite
	UnmarshalJSON(data []byte) error
}

// Point represents a coordinate pair.
type Point struct {
	X int
	Y int
}

// Ints returns the point as a pair of ints.
func (p *Point) Ints() (x, y int) {
	return p.X, p.Y
}

// Pos represents the position of an object.
type Pos struct {
	Point
	Map string
}

func NewPos(x, y int, Map string) *Pos {
	return &Pos{Point{x, y}, Map}
}

// Ints
func (p *Pos) Ints() (x, y int, Map string) {
	return p.X, p.Y, p.Map
}

type Server interface {
	Conns() map[string]*Conn
	HandleRequests()
	Maps() map[string]Map
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
}

var (
	EmptyTile Tile = Tile{Sprite(termbox.Cell{' ', termbox.ColorWhite, termbox.ColorBlack})}
	FloorTile      = Tile{Sprite(termbox.Cell{'.', termbox.ColorWhite, termbox.ColorBlack})}
)

func (t Tile) Cell() termbox.Cell {
	return termbox.Cell(t.Sprite)
}

type UI interface {
	Run()
}

// Package gorogue is a simple roguelike engine build in golang.
//
// Goals For This Project:
//
// 1. Keep it simple, stupid
//
// I'm building an engine, not a game. The idea
// is to create a simple tool that any designer can extend to create their game.
// I'm not going to bloat the repository with unnecessary things like lots of items,
// or damage equations or things like that. Users will have to add those features
// per their needs.
//
// 2. Focus on versatility
//
// I want this project to be flexible. That's why I'll be supporting a wide variety
// of  "play styles" including: Turn Based, "Real time", Multi-player, and
// Squad/Party based.
//
// 3. An engine is only as good as its documentation.
//
// 4. Implement as little as possible in the base package.
//
// Use the base package as a skeleton for any game to work from. Implement things
// in subpackages so they can just as easily be extended as removed.
package gorogue

import (
	termbox "github.com/nsf/termbox-go"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
)

func CatchSignals() {
	c := make(chan os.Signal, 2)
	signal.Notify(c)
	go func() {
		for sig := range c {
			log.Println(sig)
			switch sig {
			case syscall.SIGTERM:
				fallthrough
			case syscall.SIGKILL:
				fallthrough
			case syscall.SIGINT:
				log.Println("Exiting...")
				os.Exit(1)
			}
		}
	}()
}

// Actor is an object that can act freely. There are two main kinds of actors:
// player characters and non-player characters (NPCs). The important
// distinction being that NPCs are controlled by the server, and Player
// characters are controlled by clients.
//
// Each NPC gets their own goroutine, meaning each acts on their own thread,
// separate from other actors. The server receives requests to act from each NPC,
// and determines whether that action is valid. If if isn't, the action is rejected
// and the Actor must choose a different action to perform. If the action is valid,
// it gets stored in memory and is called during the next Map tick. (See Map.Tick)
type Actor interface {
	Object             // The Object interface.
	Move(pos Pos) bool // Moves the Actor to the given position.
}

// Command is a string mapped to an Action. Commands are  intended to be called
// in a Vi style command bar.
//
// TODO: (7) Implement Vi command bar.
type Command string

type Client interface {
	Addr() string
	Init() *UI
	Ping()
	HandleAction(a Action)
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
	Sprite() termbox.Cell
	UnmarshalJSON(data []byte) error
}

// Key is a keyboard key mapped to an Action. See package github.com/nsf/termbox-go
// for more information.
type Key struct {
	Mod termbox.Modifier // One of termbox.Mod* constants or 0.
	Key termbox.Key      // One of termbox.Key* constants, invalid if 'Ch' is not 0.
	Ch  rune             // a unicode character.
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
	// Maps() map[string]*Map
	// Ping(addr *string, reply *Pong)
	Port() string
	SetPort(port string)
}

type Tile struct {
	Sprite termbox.Cell
}

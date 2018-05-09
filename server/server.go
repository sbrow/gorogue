// Package server handles the bulk of game and NPC logic.
//
// TODO: SpawnNPC() should always trigger a new goroutine to handle that NPCs
// logic
package server

import (
	"fmt"
	. "github.com/sbrow/gorogue"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Game logic is handled on the server.
//
// Each world gets at least one goroutine, with each active map getting its own
// goroutine as well.
type Server struct {
	Host string
	Port string
	Maps []*Map
}

func Start(host, port string, maps ...*Map) *Server {
	s := &Server{host, port, maps}
	go s.HandleRequests()
	return s
}

func (s *Server) HandleRequests() {
	server := rpc.NewServer()
	server.Register(s)

	l, err := net.Listen("tcp", fmt.Sprint(s.Port))
	if err != nil {
		panic(err)
	}

	for {
		if conn, err := l.Accept(); err != nil {
			panic(err)
		} else {
			log.Println("Connection established!")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}

func (s *Server) Map(args *string, reply *Map) error {
	for _, m := range s.Maps {
		if m.Name == *args {
			*reply = *m
			break
		}
	}
	return nil
}

// TODO: Handle failure.
func (s *Server) Move(args *MoveArgs, reply *bool) error {
	// TODO: Fix
	s.Maps[0].Actors[0].SetPos(args.Points[0])
	*reply = true
	return nil
}

func (s *Server) Spawn(args Actors, reply *SpawnReply) error {
	// TODO: Fix
	m := s.Maps[0]
	args[0].SetPos(Point{5, 5})
	m.Actors = append(m.Actors, args...)
	*reply = SpawnReply{&m.Name, m.Actors}
	return nil
}

// InitiativeMode determines how Actors are given priority.
//
// In Single mode, Only one character may have priority at any given time.
//
// In Team mode, all characters on a team act  in unison.
//
// In All mode, all characters are given "identical" priority. This is used for
// realtime play.
type InitiativeMode uint8

const (
	Single InitiativeMode = iota
	Team
	All
)

// TickMode determines how Tick() is handled.
//
// In Action mode, each character can perform one action per Tick. This is the
// default mode.
//
// In AP mode, characters can perform actions whenever they have priority,
// only limited by whatever the designer wants to use instead of ticks,
// (Usually Action Points, or similar).
type TickMode uint8

const (
	Action TickMode = iota
	AP
)

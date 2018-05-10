// Package server handles the bulk of game and NPC logic.
//
// TODO: SpawnNPC() should always trigger a new goroutine to handle that NPCs
// logic
package server

import (
	"encoding/json"
	"fmt"
	. "github.com/sbrow/gorogue"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/signal"
	"syscall"
)

// Game logic is handled on the server.
//
// Each world gets at least one goroutine, with each active map getting its own
// goroutine as well.
type Server struct {
	Host string
	Port string
	Maps []*Map
	// Conns []*net.Conn
}

func Start(host, port string, maps ...*Map) *Server {
	s := &Server{host, port, maps}
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
	s.handleRequests()
	return s
}

func (s *Server) handleRequests() {
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
			log.Printf("Connection established! %s", conn.RemoteAddr())
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

func (s *Server) Move(args *Action, reply *ActionResponse) error {
	m := s.Maps[0]
	for _, a := range m.Actors() {
		if a.Name() == args.Caller {
			_ = m.WaitForTurn(a)
			var p *Pos
			_ = json.Unmarshal(args.Args[0], &p)
			a.Move(*p)
			*reply = ActionResponse{
				Msg:   "Success",
				Reply: true,
			}
			return nil
		}
	}
	*reply = ActionResponse{
		Msg:   "Actor not found.",
		Reply: false,
	}
	return nil
}

// Spawn spawns new actors on the first map.
func (s *Server) Spawn(args Actors, reply *SpawnReply) error {
	Map := 0
	m := s.Maps[Map]
	args[0].SetPos(*NewPos(5, 5, Map))
	for _, a := range args {
		switch v := a.(type) {
		case Player:
			m.Players = append(m.Players, v)
		}
	}
	*reply = SpawnReply{&m.Name, args} // TODO: Fix
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
	ActionBased TickMode = iota
	AP
)

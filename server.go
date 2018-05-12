// TODO: SpawnNPC() should always trigger a new goroutine to handle that NPCs
package gorogue

import (
	"fmt"
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
	Port  string
	Maps  map[string]*Map
	Conns map[string]*Conn
}

type Conn struct {
	Conn  *net.Conn
	Squad Actors
}

func NewServer(port string, maps ...*Map) *Server {
	s := &Server{
		Port:  port,
		Maps:  map[string]*Map{},
		Conns: map[string]*Conn{},
	}
	for _, m := range maps {
		s.Maps[m.Name] = m
	}
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
	log.Println("Waiting for players...")
	for {
		if conn, err := l.Accept(); err != nil {
			panic(err)
		} else {
			log.Printf("Connection established on %s", conn.RemoteAddr())
			addr := fmt.Sprint(conn.RemoteAddr())
			conn.Write([]byte(addr))
			s.Conns[addr] = &Conn{Conn: &conn, Squad: []Actor{}}
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}

type Pong struct {
	Squad Actors
	Maps  map[string]*Map
}

func (s *Server) Ping(addr *string, reply *Pong) error {
	pong := &Pong{}
	// TODO: (10) Reduce Pong to only relevant maps.
	pong.Maps = s.Maps
	for _, p := range s.Conns[*addr].Squad {
		pong.Squad = append(pong.Squad, p)
	}

	*reply = *pong
	return nil
}

func (s *Server) Move(args *MoveAction, reply *ActionResponse) error {
	log.Println("Recieved action", *args)
	// TODO: (2) Temporary map Fix
	var m *Map
	for _, m = range s.Maps {
		break
	}
	for _, p := range m.Players {
		if p.Name() == args.Caller {
			_ = m.WaitForTurn(p)
			p.Move(args.Pos)

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
func (s *Server) Spawn(args *SpawnAction, reply *bool) error {
	fmt.Println("Args", *args)
	// TODO: (2) Temporary map Fix
	var Map string
	for Map, _ = range s.Maps {
		break
	}
	m := s.Maps[Map]
	sq := s.Conns[args.Caller].Squad
	for _, a := range args.Actors {
		switch v := a.(type) {
		case Player:
			v.SetPos(Pos{Point{5, 5}, Map})
			_, present := m.Players[v.ID()]
			for present {
				v.SetIndex(v.Index() + 1)
				_, present = m.Players[v.ID()]
			}
			m.Players[a.ID()] = v
			sq = append(sq, m.Players[v.ID()].(Player))
		}
	}
	s.Conns[args.Caller].Squad = sq
	fmt.Println("Caller:", args.Caller, "Actors", args.Actors[0])
	fmt.Printf("%+v\n", s.Conns[args.Caller])
	*reply = true
	go m.Tick()
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
type initiativeMode uint8

const (
	single initiativeMode = iota
	team
	all
)

// TickMode determines how Tick() is handled.
//
// In Action mode, each character can perform one action per Tick. This is the
// default mode.
//
// In AP mode, characters can perform actions whenever they have priority,
// only limited by whatever the designer wants to use instead of ticks,
// (Usually Action Points, or similar).
type tickMode uint8

const (
	actionBased tickMode = iota
	ap
)

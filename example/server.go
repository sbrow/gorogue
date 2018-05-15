// TODO: SpawnNPC() should always trigger a new goroutine to handle that NPCs
package example

import (
	"fmt"
	engine "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/action"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Pong struct {
	Squad Actors
	Maps  map[string]*Map
}

// Game logic is handled on the server.
//
// Each world gets at least one goroutine, with each active map getting its own
// goroutine as well.
type Server struct {
	engine.Server
	port  string
	maps  map[string]*Map
	conns map[string]*engine.Conn
}

func (s *Server) Conns() map[string]*engine.Conn {
	return s.conns
}

func (s *Server) Disconnect(args *string, reply *string) error {
	log.Printf("%s Disconnected.\n", *args)
	for _, actor := range s.conns[*args].Squad {
		for _, v := range s.maps {
			if v.Remove(actor) == true {
				break
			}
		}
	}
	s.conns[*args] = nil
	return nil
}

func (s *Server) HandleRequests() {
	if s.maps == nil {
		s.maps = map[string]*Map{
			"Map_1": NewMap(24, 24, "Map_1"),
		}
	}
	if s.conns == nil {
		s.conns = map[string]*engine.Conn{}
	}
	server := rpc.NewServer()
	server.Register(s)

	l, err := net.Listen("tcp", fmt.Sprint(s.port))
	if err != nil {
		panic(err)
	}
	log.Println("Waiting for players...")
	for {
		if conn, err := l.Accept(); err != nil {
			panic(err)
		} else {
			log.Printf("%s Connected.", conn.RemoteAddr())
			addr := fmt.Sprint(conn.RemoteAddr())
			conn.Write([]byte(addr))
			s.conns[addr] = &engine.Conn{Conn: &conn, Squad: []engine.Actor{}}
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}

func (s *Server) Maps() map[string]engine.Map {
	var out map[string]engine.Map
	for k, v := range s.maps {
		out[k] = v
	}
	return out
}

func (s *Server) Move(args *action.Move, reply *string) error {
	log.Println("Recieved action", args.String())
	// TODO: (2) Temporary map Fix
	var m *Map
	for _, m = range s.maps {
		break
	}
	for _, p := range m.Players {
		if p.Name() == args.Caller {
			_ = m.WaitForTurn(1)
			p.Move(args.Pos)
			return nil
		}
	}
	return nil
}

func (s *Server) Ping(addr *string, reply *Pong) error {
	pong := &Pong{}
	// TODO: (10) Reduce Pong to only relevant maps.
	pong.Maps = s.maps
	for _, p := range s.conns[*addr].Squad {
		pong.Squad = append(pong.Squad, p)
	}
	*reply = *pong
	return nil
}

func (s *Server) SetPort(port string) {
	s.port = port
}

// Spawn spawns new actors on the first map.
func (s *Server) Spawn(args *Spawn, reply *bool) error {
	// TODO: (2) Temporary map Fix
	var Map string
	for Map, _ = range s.maps {
		break
	}
	m := s.maps[Map]
	sq := s.conns[args.Caller].Squad
	for _, a := range args.Actors {
		switch v := a.(type) {
		case *Player:
			v.SetPos(engine.Pos{engine.Point{5, 5}, Map})
			m.Players = append(m.Players, v)
			sq = append(sq, v)
		}
	}
	s.conns[args.Caller].Squad = sq
	*reply = true
	go m.Tick()
	return nil
}

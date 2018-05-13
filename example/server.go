// TODO: SpawnNPC() should always trigger a new goroutine to handle that NPCs
package example

import (
	"fmt"
	engine "github.com/sbrow/gorogue"
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
	Maps  map[string]*Map
	conns map[string]*engine.Conn
}

func (s *Server) Conns() map[string]*engine.Conn {
	return s.conns
}

func (s *Server) HandleRequests() {
	if s.Maps == nil {
		s.Maps = map[string]*Map{
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
			log.Printf("Connection established on %s", conn.RemoteAddr())
			addr := fmt.Sprint(conn.RemoteAddr())
			conn.Write([]byte(addr))
			s.conns[addr] = &engine.Conn{Conn: &conn, Squad: []engine.Actor{}}
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}

func (s *Server) Ping(addr *string, reply *Pong) error {
	pong := &Pong{}
	// TODO: (10) Reduce Pong to only relevant maps.
	pong.Maps = s.Maps
	for _, p := range s.conns[*addr].Squad {
		pong.Squad = append(pong.Squad, p)
	}
	*reply = *pong
	return nil
}

func (s *Server) Move(args *MoveAction, reply *engine.ActionResponse) error {
	log.Println("Recieved action", *args)
	// TODO: (2) Temporary map Fix
	var m *Map
	for _, m = range s.Maps {
		break
	}
	for _, p := range m.Players {
		if p.Name() == args.Caller {
			_ = m.WaitForTurn(1)
			p.Move(args.Pos)
			*reply = engine.ActionResponse{
				Reply: true,
			}
			return nil
		}
	}
	*reply = engine.ActionResponse{
		Reply: false,
	}
	*reply.Msg = "Actor not found."
	return nil
}

func (s *Server) SetPort(port string) {
	s.port = port
}

// Spawn spawns new actors on the first map.
func (s *Server) Spawn(args *SpawnAction, reply *bool) error {
	log.Println("Args", *args)
	// TODO: (2) Temporary map Fix
	var Map string
	for Map, _ = range s.Maps {
		break
	}
	m := s.Maps[Map]
	sq := s.conns[args.Caller].Squad
	for _, a := range args.Actors {
		switch v := a.(type) {
		case *Player:
			v.SetPos(engine.Pos{engine.Point{5, 5}, Map})
			m.Players = append(m.Players, v)
			sq = append(sq, v)
			/*
				_, present := m.Players[v.ID()]
				for present {
					v.SetIndex(v.Index() + 1)
					_, present = m.Players[v.ID()]
				}
				m.Players[a.ID()] = v
				sq = append(sq, m.Players[v.ID()].(Player))
			*/
		}
	}
	s.conns[args.Caller].Squad = sq
	fmt.Println("Caller:", args.Caller, "Actor", args.Actors)
	fmt.Printf("%+v\n", s.conns[args.Caller])
	*reply = true
	go m.Tick()
	return nil
}

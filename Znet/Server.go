package Znet

import (
	"fmt"
	"net"
	"zinx/Zinterface"
)

type Server struct {
	name      string
	ipVersion string
	ip        string
	port      int
	msgHandle Zinterface.IMsgHandle
	cmgr      Zinterface.IConnManger
}

// func callback(c *net.TCPConn, buf []byte, cnt int) error {
// 	fmt.Printf("Connection HandleFuc Start...\n")
// 	if _, err := c.Write(buf[:cnt]); err != nil {
// 		fmt.Printf("Connection Writer Error : %s\n", err)
// 		return err
// 	}
// 	return nil
// }
func (s *Server) Start() {
	go func() {
		s.msgHandle.StartWorkerPool()
		fmt.Println("Listenning Server ", s.name, "    ", s.ip, ":", s.port, "............")
		addr, err := net.ResolveTCPAddr(s.ipVersion, fmt.Sprintf("%s:%d", s.ip, s.port))
		if err != nil {
			fmt.Printf("ResolveTCP Failed!\n")
			return
		}
		listener, err := net.ListenTCP(s.ipVersion, addr)
		if err != nil {
			fmt.Printf("ListenTCP Error : %s\n", err)
			return
		}
		var cid uint32 = 0

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf(" AcceptTCP Error: %s\n", err)
				return
			}
			if s.cmgr.ConnSize() > 3 {
				fmt.Println(" connection num is maxium .....")
				conn.Close()
				continue
			}
			c := NewConnection(s, conn, cid, s.msgHandle)

			cid++
			go c.Start()
			// go func(){
			// 	for{
			// buf := make([]byte, 512)
			// cnt,err := conn.Read(buf)
			// if err != nil{
			// 	fmt.Printf("Read %d byte Error", cnt)
			// 	continue
			// }
			// if _,err := conn.Write(buf[:cnt]);err !=nil{
			// 	fmt.Printf("Write %d byte Error",cnt)
			// 	continue
			// }
			// 	}
			// }()

		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("Server is Stoped")
	s.cmgr.Clear()
}

func (s *Server) Run() {
	s.Start()

	select {}
}

func (s *Server) AddRouter(msgid uint32, r Zinterface.IRouter) {
	s.msgHandle.AddRouter(msgid, r)
	fmt.Printf("Add router succ...\n")
}

func (s *Server) GetConnMgr() Zinterface.IConnManger {
	return s.cmgr
}
func NewServer(name, ipv, ip string, port int) *Server {
	s := &Server{
		name:      name,
		ipVersion: ipv,
		ip:        ip,
		port:      port,
		msgHandle: NewMsgHandle(),
		cmgr:      NewConnManger(),
	}
	return s
}

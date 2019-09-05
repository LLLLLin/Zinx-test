package Znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/Zinterface"
)

type Connection struct {
	server  Zinterface.IServer
	conn    *net.TCPConn
	conId   uint32
	isClose bool
	//handleAPI Zinterface.HandleFunc
	exitChan  chan bool
	datachan  chan []byte
	msgHandle Zinterface.IMsgHandle
}

func (c *Connection) StartReader() {
	fmt.Printf("[Read goroutine start...]\n")
	defer c.Stop()
	defer fmt.Println("[Read goroutine is stoped...]")
	for {
		// 	buf := make([]byte, 512)
		// 	_, err := c.conn.Read(buf)
		// 	if err != nil {
		// 		fmt.Printf("Connection Read Error %s\n", err)
		// 		continue
		// 	}
		// if err := c.handleAPI(c.conn, buf, cnt); err != nil {
		// 	fmt.Printf("Connection HandleAPI Error %s\n", err)
		// 	continue
		// }
		dp := NewDatapack()
		msghead := make([]byte, dp.GetMsgHeadLen())
		_, err := io.ReadFull(c.GetTcpCon(), msghead)
		if err != nil {
			fmt.Println("Connection Read msghead err", err)
			return
		}
		msg, _ := dp.Unpack(msghead)
		if msg.GetMsgDataLen() > 0 {
			msgdata := make([]byte, msg.GetMsgDataLen())
			_, err := io.ReadFull(c.GetTcpCon(), msgdata)
			if err != nil {
				fmt.Println("Connection getmsgdata error", err)
				return
			}
			msg.SetMsgData(msgdata)
		}
		fmt.Println("[recev msg id]:", msg.GetMsgId(),
			"[msg len]:", msg.GetMsgDataLen(),
			"[msg data]:", string(msg.GetMsgData()))

		r := Request{
			conn: c,
			msg:  msg,
		}

		// go c.msgHandle.DoMsgHandle(&r)
		c.msgHandle.SendMsgToQueue(&r)
	}
}

func (c *Connection) StartWrite() {
	fmt.Println("[Write groutine is started ... ]")
	defer fmt.Println("[Write groutine is stoped ... ]")
	for {
		select {
		case data := <-c.datachan:
			if _, err := c.GetTcpCon().Write(data); err != nil {
				fmt.Println(" Connection write error", err)
				continue
			}
		case <-c.exitChan:
			return
		}
	}
}
func (c *Connection) Start() {
	fmt.Printf("Connection id: %d is starting\n", c.conId)
	go c.StartReader()
	go c.StartWrite()

}
func (c *Connection) Stop() {
	fmt.Printf("Connection id: %d  Stop\n", c.conId)
	if c.isClose == true {
		return
	}
	c.isClose = true
	c.exitChan <- true

	c.conn.Close()
	c.server.GetConnMgr().Remove(c)
	close(c.datachan)
	close(c.exitChan)
}
func (c *Connection) GetId() uint32 {
	return c.conId
}
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) SendMsg(id uint32, data []byte) error {
	if c.isClose == true {
		return errors.New("Connection is closed...")
	}
	dp := NewDatapack()
	binaryMsg, err := dp.Pack(NewMessage(id, data))
	if err != nil {
		fmt.Println(" Connection pack error", err)
		return err
	}

	c.datachan <- binaryMsg
	return nil

}
func (c *Connection) GetTcpCon() *net.TCPConn {
	return c.conn
}

func NewConnection(s Zinterface.IServer, c *net.TCPConn, id uint32, handle Zinterface.IMsgHandle) *Connection {
	con := &Connection{
		server: s,
		conn:   c,
		conId:  id,
		//handleAPI: callback,
		isClose:   false,
		exitChan:  make(chan bool),
		datachan:  make(chan []byte),
		msgHandle: handle,
	}
	s.GetConnMgr().Add(con)
	return con
}

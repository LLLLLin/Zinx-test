package Zinterface

import (
	"net"
)

type IConnection interface {
	Start()
	Stop()
	GetId() uint32
	GetRemoteAddr() net.Addr
	GetTcpCon() *net.TCPConn
	SendMsg(uint32, []byte) error
}

//type HandleFunc func(*net.TCPConn, []byte, int) error

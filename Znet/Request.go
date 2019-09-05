package Znet

import (
	"zinx/Zinterface"
)

type Request struct {
	conn Zinterface.IConnection
	msg  Zinterface.IMessage
}

func (r *Request) GetConnection() Zinterface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetMsgData()
}
func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}

func NewRequest(c Zinterface.IConnection, m Zinterface.IMessage) *Request {
	r := &Request{
		conn: c,
		msg:  m,
	}
	return r
}

package Zinterface

type IDatapack interface {
	GetMsgHeadLen() uint32
	Pack(IMessage) ([]byte, error)
	Unpack([]byte) (IMessage, error)
}

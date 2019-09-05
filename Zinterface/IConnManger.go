package Zinterface

type IConnManger interface {
	Add(IConnection)
	Remove(IConnection)
	Get(uint32) (IConnection, error)
	ConnSize() int
	Clear()
}

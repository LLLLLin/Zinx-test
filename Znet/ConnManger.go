package Znet

import (
	"errors"
	"fmt"
	"sync"
	"zinx/Zinterface"
)

type ConnManger struct {
	conns    map[uint32]Zinterface.IConnection
	connLock sync.RWMutex
}

func (cm *ConnManger) Add(c Zinterface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	cm.conns[c.GetId()] = c
	fmt.Println("connection add success , conn num:", cm.ConnSize())
}

func (cm *ConnManger) Remove(c Zinterface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	if _, ok := cm.conns[c.GetId()]; ok != true {
		fmt.Println("No exist conn connid:", c.GetId())
		return
	}
	delete(cm.conns, c.GetId())

	fmt.Println("connection delete success , conn num:", cm.ConnSize())
}
func (cm *ConnManger) Get(cid uint32) (Zinterface.IConnection, error) {
	cm.connLock.RLocker()
	defer cm.connLock.RUnlock()
	if c, ok := cm.conns[cid]; ok == true {
		return c, nil
	} else {
		return nil, errors.New("Not Found connecion id: " + string(cid))
	}
}

func (cm *ConnManger) ConnSize() int {
	return len(cm.conns)

}

func (cm *ConnManger) Clear() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	for cid, _ := range cm.conns {
		cm.conns[cid].Stop()
		delete(cm.conns, cid)
	}
	fmt.Println("Clear all connection conn num:", cm.ConnSize())
}

func NewConnManger() *ConnManger {
	cm := &ConnManger{
		conns: make(map[uint32]Zinterface.IConnection),
	}
	return cm
}

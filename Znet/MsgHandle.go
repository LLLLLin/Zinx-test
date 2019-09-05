package Znet

import (
	"fmt"
	"zinx/Zinterface"
)

type MsgHandle struct {
	apis         map[uint32]Zinterface.IRouter
	workPoolSize uint32
	taskqueue    []chan Zinterface.IRequest
}

func (mh *MsgHandle) DoMsgHandle(req Zinterface.IRequest) {
	if _, ok := mh.apis[req.GetMsgId()]; ok != true {
		fmt.Println("No exist this router")
		return
	}
	mh.apis[req.GetMsgId()].PreHandle(req)
	mh.apis[req.GetMsgId()].Handle(req)
	mh.apis[req.GetMsgId()].PostHandle(req)
}

func (mh *MsgHandle) AddRouter(msgid uint32, router Zinterface.IRouter) {
	if _, ok := mh.apis[msgid]; ok == true {
		println("the msg has been already router")
		return
	}
	mh.apis[msgid] = router
	println("msg id :", msgid, " add router success...")
}

func (mh *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(mh.workPoolSize); i++ {
		mh.taskqueue[i] = make(chan Zinterface.IRequest, 1024)
		go mh.StratWorker(i, mh.taskqueue[i])
	}
}

func (mh *MsgHandle) StratWorker(workerId int, tq chan Zinterface.IRequest) {
	fmt.Println("Worker Id :", workerId, "is started")
	for {
		select {
		case req := <-tq:
			mh.DoMsgHandle(req)
		}
	}
}

func (mh *MsgHandle) SendMsgToQueue(req Zinterface.IRequest) {
	workerId := req.GetConnection().GetId() % mh.workPoolSize
	fmt.Println("Send ConnectionID :", req.GetConnection().GetId(),
		"request MsgID :", req.GetMsgId(),
		"to workerID :", workerId)
	mh.taskqueue[workerId] <- req
}
func NewMsgHandle() *MsgHandle {
	m := &MsgHandle{
		apis:         make(map[uint32]Zinterface.IRouter),
		workPoolSize: 8,
		taskqueue:    make([]chan Zinterface.IRequest, 8),
	}
	return m
}

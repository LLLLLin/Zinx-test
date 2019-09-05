package Zinterface

type IMsgHandle interface {
	DoMsgHandle(IRequest)
	AddRouter(uint32, IRouter)
	StartWorkerPool()
	SendMsgToQueue(IRequest)
}

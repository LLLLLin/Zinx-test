package Zinterface

type IServer interface {
	Start()

	Stop()

	Run()

	AddRouter(uint32, IRouter)

	GetConnMgr() IConnManger
}

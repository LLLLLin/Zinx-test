package Zinterface

type IMessage interface {
	GetMsgId() uint32

	GetMsgDataLen() uint32

	GetMsgData() []byte

	SetMsgId(uint32)

	SetMsgDataLen(uint32)

	SetMsgData([]byte)
}

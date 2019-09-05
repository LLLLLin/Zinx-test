package Znet

type Message struct {
	msgDataLen uint32
	msgId      uint32
	msgData    []byte
}

func (m *Message) GetMsgId() uint32 {
	return m.msgId
}

func (m *Message) GetMsgDataLen() uint32 {
	return m.msgDataLen
}

func (m *Message) GetMsgData() []byte {
	return m.msgData
}

func (m *Message) SetMsgId(id uint32) {
	m.msgId = id
}
func (m *Message) SetMsgDataLen(len uint32) {
	m.msgDataLen = len
}

func (m *Message) SetMsgData(data []byte) {
	m.msgData = data
}

func NewMessage(id uint32, data []byte) *Message {
	m := &Message{
		msgId:      id,
		msgDataLen: uint32(len(data)),
		msgData:    data,
	}
	return m
}

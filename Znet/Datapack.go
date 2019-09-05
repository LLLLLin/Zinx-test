package Znet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"zinx/Zinterface"
)

type Datapack struct {
}

func NewDatapack() *Datapack {
	return &Datapack{}
}
func (dp *Datapack) GetMsgHeadLen() uint32 {
	return 8
}

func (dp *Datapack) Pack(msg Zinterface.IMessage) ([]byte, error) {
	databuff := bytes.NewBuffer([]byte{})

	if err := binary.Write(databuff, binary.LittleEndian, msg.GetMsgDataLen()); err != nil {
		fmt.Println("Pack Write DataLen error", err)
		return nil, err
	}
	if err := binary.Write(databuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		fmt.Println("Pack Write Id error", err)
		return nil, err
	}
	if err := binary.Write(databuff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		fmt.Println("Pack Write Data error", err)
		return nil, err
	}
	return databuff.Bytes(), nil
}

func (dp *Datapack) Unpack(data []byte) (Zinterface.IMessage, error) {
	dataReader := bytes.NewReader(data)
	msg := &Message{}
	if err := binary.Read(dataReader, binary.LittleEndian, &msg.msgDataLen); err != nil {
		fmt.Println("UnPack Read DataLen error", err)
		return nil, err
	}

	if err := binary.Read(dataReader, binary.LittleEndian, &msg.msgId); err != nil {
		fmt.Println("UnPack Read ID error", err)
		return nil, err
	}
	return msg, nil
}

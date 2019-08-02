package protocol

import (
	"fmt"
	"io"
	"net"
	"wentby/config"
)

// PacketReader is used to unmarshal a complete packet from buff
type ReaderInter interface {
	// Read data from conn and build a complete packet.
	ReadPacket(conn net.Conn) (interface{}, error)
}

// PacketWriter is used to marshal packet into buff
type WriterInter interface {
	WritePacket(net.Conn, interface{}) error
}

// PacketProtocol just a composite interface
type ProtocolInter interface {
	ReaderInter
	WriterInter
}

type ProtocolImpl struct {
}

func (pi *ProtocolImpl) ParaseHead(packet interface{}, buff []byte) (interface{}, error) {
	if len(buff) < 16 {
		return nil, config.ErrBuffLenLess
	}
	msgpacket, ok := packet.(*MsgPacket)
	if !ok {
		fmt.Println("it's not msgpacket type")
		return nil, config.ErrTypeAssertain
	}
	stream := NewBigEndianStream(buff)
	var err error
	if msgpacket.Head.Id, err = stream.ReadUint16(); err != nil {
		return nil, config.ErrParaseMsgHead
	}

	if msgpacket.Head.Len, err = stream.ReadUint16(); err != nil {
		return nil, config.ErrParaseMsgHead
	}

	return msgpacket, nil
}

func (pi *ProtocolImpl) ReadPacket(conn net.Conn) (interface{}, error) {
	buff := make([]byte, 1024)
	_, err := io.ReadAtLeast(conn, buff[:16], 16)
	if err != nil {
		return nil, config.ErrReadAtLeast
	}

	var value interface{}
	var msgpacket *MsgPacket = new(MsgPacket)
	value, err = pi.ParaseHead(msgpacket, buff[:16])

	msgpacket, ok := value.(*MsgPacket)
	if !ok {
		fmt.Println("it's not msgpacket type")
		return nil, config.ErrTypeAssertain
	}

	if uint16(len(buff[16:])) < msgpacket.Head.Len {
		return nil, config.ErrMsgLenLarge
	}

	if _, err = io.ReadFull(conn, buff[16:msgpacket.Head.Len+16]); err != nil {
		return nil, config.ErrReadAtLeast
	}

	return msgpacket, nil
}

func (pi *ProtocolImpl) WritePacket(conn net.Conn, packet interface{}) error {
	return nil
}
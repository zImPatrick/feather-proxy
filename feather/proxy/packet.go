package proxy

import (
	"encoding/binary"
	"io"
	"net"
)

func readPacket(c net.Conn) ([]byte, error) {
	lengthBuf := make([]byte, 4)
	_, err := io.ReadFull(c, lengthBuf)
	if err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint32(lengthBuf)

	packetBuf := make([]byte, length)
	_, err = io.ReadFull(c, packetBuf)
	if err != nil {
		return nil, err
	}

	return packetBuf, nil
}

func writePacket(c net.Conn, data []byte) error {
	lengthBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBuf, uint32(len(data)))

	packet := append(lengthBuf, data...)
	_, err := c.Write(packet)
	if err != nil {
		return err
	}

	return nil
}

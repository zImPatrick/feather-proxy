package proxy

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

type CommunicationConnection struct {
	Host                 string
	ConnectionServerHost string
	Client               *FeatherProxyClient
	Token                string
	connection           net.Conn
	exitSignal           chan int
}

func (c *CommunicationConnection) Connect() error {
	c.exitSignal = make(chan int, 1)

	connection, err := net.Dial("tcp", c.Host)
	if err != nil {
		return err
	}

	c.connection = connection

	writePacket(connection, []byte(c.Token))

	data, err := readPacket(connection)
	if err != nil {
		return err
	}

	// check if it's the ok packet
	// 0x4F, 0x4B is "OK"
	if !bytes.Equal([]byte{0x4f, 0x4b}, data) {
		return fmt.Errorf("expected ok packet, got %x against %x", data, []byte{0x4F, 0x4B})
	}

	go c.readPackets()
	go c.heartbeat()

	return nil
}

func (c *CommunicationConnection) heartbeat() error {
	for {
		_, err := c.connection.Write([]byte{0x00})
		if err != nil {
			return err
		}
		time.Sleep(60_000)
	}
}

func (c *CommunicationConnection) readPackets() error {
	for {
		data, err := readPacket(c.connection)
		if err != nil {
			c.connection.Close()
			return err
		}

		switch len(data) {
		case 16: // is probably a join request
			conn := ConnectionConnection{
				Host:   c.ConnectionServerHost,
				Client: c.Client,
			}

			err := conn.Connect(data)
			if err != nil {
				c.Client.Logger.Printf("An error occured while connecting to the connection server: %s", err.Error())
			}
		default:
			c.Client.Logger.Printf("Unknown packet on communication connection with length %d: %x", len(data), data)
		}
	}
}

func (c *CommunicationConnection) Disconnect() error {
	return c.connection.Close()
}

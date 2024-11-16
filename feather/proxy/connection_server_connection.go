package proxy

import (
	"fmt"
	"net"
)

// this might need a rename
type ConnectionConnection struct {
	Client *FeatherProxyClient
	Host   string
}

func (c *ConnectionConnection) Connect(token []byte) error {
	connection, err := net.Dial("tcp", c.Host)
	if err != nil {
		return err
	}

	_, err = connection.Write(token)
	if err != nil {
		return err
	}

	fmt.Println("Giving connection to pendingConnections channel")
	c.Client.pendingConnections <- &connection
	return nil
}

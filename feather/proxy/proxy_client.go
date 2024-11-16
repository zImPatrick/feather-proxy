package proxy

import (
	"log"
	"net"
)

type FeatherProxyClient struct {
	ServerToken              string
	communicationConnections []CommunicationConnection
	pendingConnections       chan *net.Conn
}

func (c *FeatherProxyClient) Connect() error {
	c.pendingConnections = make(chan *net.Conn, 16)

	proxies, err := retrieveProxyList()
	if err != nil {
		return err
	}

	c.communicationConnections = make([]CommunicationConnection, len(proxies))
	for i, proxy := range proxies {
		connection := CommunicationConnection{
			Host:                 proxy.CommunicationHost,
			ConnectionServerHost: proxy.ConnectionHost,
			Token:                c.ServerToken,
			Client:               c,
		}

		err := connection.Connect()
		if err != nil {
			log.Printf("An error occured while connecting to communication host %s: %s", proxy.CommunicationHost, err.Error())
		}
		c.communicationConnections[i] = connection
	}

	return nil
}

// Implement net.Listener
func (c *FeatherProxyClient) Accept() (net.Conn, error) {
	conn := <-c.pendingConnections
	return *conn, nil
}

func (c *FeatherProxyClient) Close() error {
	for _, proxy := range c.communicationConnections {
		proxy.Disconnect()
	}

	return nil
}

// return a fake ip
func (c *FeatherProxyClient) Addr() net.Addr {
	ip, _ := net.ResolveIPAddr("ip", "127.0.0.1")
	return ip
}

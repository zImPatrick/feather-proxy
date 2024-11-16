package proxy

import (
	"log"
	"net"
)

type FeatherProxyClient struct {
	ServerToken              string
	communicationConnections []CommunicationConnection
	pendingConnections       chan *net.Conn
	Logger                   *log.Logger
}

// Connects to all proxy endpoints (as this is what the official client also does)
func (c *FeatherProxyClient) Connect() error {
	// todo: There should probably be a check here to make sure
	// that we aren't already connected

	// We should probably move this into some sort of New() function
	// that just returns an instance of FeatherProxyClient with
	// logger and pendingConnections set
	if c.Logger == nil {
		c.Logger = log.Default()
	}

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

		// todo: this could be improved by connecting in parallel?
		// the official client also does this but I'm not doing it
		// here for simplicity reasons.. for now
		err := connection.Connect()
		if err != nil {
			c.Logger.Printf("An error occured while connecting to communication host %s: %s", proxy.CommunicationHost, err.Error())
		}
		c.communicationConnections[i] = connection
	}

	return nil
}

// The next few functions (Accept, Close, Addr) are only here
// to implement net.Listener

// Waits for a connection and accepts it
func (c *FeatherProxyClient) Accept() (net.Conn, error) {
	conn := <-c.pendingConnections
	return *conn, nil
}

// Disconnects from all proxy servers
func (c *FeatherProxyClient) Close() error {
	for _, proxy := range c.communicationConnections {
		proxy.Disconnect()
	}

	return nil
}

// Should return listening ip, but this isn't really
// applicable here so it just returns 127.0.0.1
func (c *FeatherProxyClient) Addr() net.Addr {
	ip, _ := net.ResolveIPAddr("ip", "127.0.0.1")
	return ip
}

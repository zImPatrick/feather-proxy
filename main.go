package main

import (
	"feather-proxy/feather/proxy"
	"flag"
	"io"
	"log"
	"net"
)

func getConfig(Path string) Config {
	conf := Config{}
	err := conf.Load(Path)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	return conf
}

func main() {
	// get args
	configPath := flag.String("config_path", "config.json", "tells the application where to look for the config")
	serverToken := flag.String("server_token", "", "specifies the server token, defaults to the one in the config")
	proxyTarget := flag.String("target", "127.0.0.1:25565", "where the application should proxy traffic to")

	flag.Parse()

	// get config
	config := getConfig(*configPath)
	if *serverToken == "" {
		serverToken = &config.ServerToken
	}

	// error out if we don't have any server token
	if *serverToken == "" {
		log.Fatalf("No server token was provided")
	}

	client := proxy.FeatherProxyClient{
		ServerToken: *serverToken,
	}

	log.Println("Connecting to Feather Client's proxies...")
	err := client.Connect()
	if err != nil {
		panic(err)
	}

	log.Println("Ready for connections")

	for {
		// This is something for the library, but there are currently
		// no errors that are thrown on Accept. This is something very
		// worth looking into
		conn, err := client.Accept()
		if err != nil {
			log.Printf("Failed accepting a connection: %v", err)
			continue
		}

		proxyConnection, err := net.Dial("tcp", *proxyTarget)
		if err != nil {
			log.Printf("Failed connecting to proxy target: %v", err)
			continue
		}

		go io.Copy(conn, proxyConnection)
		go io.Copy(proxyConnection, conn)
	}
}

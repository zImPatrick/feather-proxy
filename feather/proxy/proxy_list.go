package proxy

import (
	"encoding/json"
	"io"
	"net/http"
)

type Proxy struct {
	CommunicationHost string `json:"server_communication_hostname"`
	ConnectionHost    string `json:"server_connection_hostname"`
}

func retrieveProxyList() ([]Proxy, error) {
	resp, err := http.Get("https://launcher-client.feathermc.com/proxy_list.json")
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var proxyList []Proxy
	err = json.Unmarshal(data, &proxyList)
	if err != nil {
		return nil, err
	}

	return proxyList, nil
}

# Proxy protocol

## Getting a server token
To get a server token, you have to send a "Create Server" request to Feather's API. This is mentioned in the documentation about [Player Servers in general](./player-servers.md).

## Getting a list of proxy servers
Feather currently sends a request to `https://launcher-client.feathermc.com/proxy_list.json` to get a list of their proxies.

It looks like this:
```json
[
  {
    "server_communication_hostname": "de3.proxy.feathermc.net:2343",
    "server_connection_hostname": "de3.proxy.feathermc.net:2344"
  },
  ...
]
```

## How it works
When a server is started, the client connects to all(?) communication servers and [authenticates](#authentication) to every one of them. If authentication is successful, they will just respond with "OK". The client then starts heartbeating in (currently unknown) intervals. Heartbeats are just packets with zero data.

## Packet format
A packet consists of just two parts: length and data.
The first part, length, is just a 4-byte big-endian (probably unsigned) number. It tells us how long the data segment is.
The data segment are just bytes.

Example of an OK packet:
```
0000 0002 	4f4b	....OK
^^^^    	^^^^
length		data
```

## Communication server

### Authentication
This is the first packet that gets sent. It gets sent by the client and only contains the server token in text.

<details>
	<summary>Example</summary>
	
	```
	0000 017f 6579 4a68 6247 6369 4f69 4a49  ....eyJhbGciOiJI
	557a 4931 4e69 4973 496e 5235 6343 4936  UzI1NiIsInR5cCI6
	496b 7058 5643 4a39 2e65 794a 3163 3256  IkpXVCJ9.eyJ1c2V
	7953 5551 694f 6949 7959 3251 784e 474d  ySUQiOiIyY2QxNGM
	784e 5331 6859 574a 694c 5451 7959 6a45  xNS1hYWJiLTQyYjE
	7459 546c 6a4e 4331 6959 6d4d 314e 6a6c  tYTljNC1iYmM1Njl
	6d4d 6a45 775a 6a6b 694c 434a 7261 5735  mMjEwZjkiLCJraW5
	6b49 6a6f 6963 4778 6865 5756 7949 6977  kIjoicGxheWVyIiw
	6964 4739 725a 5735 4c61 5735 6b49 6a6f  idG9rZW5LaW5kIjo
	6963 4778 6865 5756 7958 334e 6c63 6e5a  icGxheWVyX3NlcnZ
	6c63 6949 7349 6e52 7661 3256 7552 4756  lciIsInRva2VuRGV
	3059 576c 7363 7949 3665 794a 3062 3274  0YWlscyI6eyJ0b2t
	6c62 6b6c 4549 6a6f 694d 6d4d 344e 5451  lbklEIjoiMmM4NTQ
	794d 7a41 744d 6a4d 344e 4330 304d 7a64  yMzAtMjM4NC00Mzd
	694c 546b 7959 6a67 744e 324e 6c5a 6a41  iLTkyYjgtN2NlZjA
	785a 6d46 6a59 7a5a 6d49 6977 6963 3256  xZmFjYzZmIiwic2V
	7964 6d56 7953 5551 694f 6949 7a4f 4463  ydmVySUQiOiIzODc
	7959 6a4d 315a 6930 774e 445a 694c 5451  yYjM1Zi0wNDZiLTQ
	794d 6d49 7459 6a63 325a 4331 6d59 6d45  yMmItYjc2ZC1mYmE
	314e 4752 6959 5755 315a 444d 6966 5377  1NGRiYWU1ZDMifSw
	6961 5746 3049 6a6f 784e 7a49 784d 5459  iaWF0IjoxNzIxMTY
	324d 5441 7a66 512e 6932 5966 4f36 644c  2MTAzfQ.i2YfO6dL
	3339 7844 3967 4f52 6f6a 7244 364b 7156  39xD9gORojrD6KqV
	4a2d 6c5a 5739 6c7a 3755 4569 4554 7862  J-lZW9lz7UEiETxb
	5f76 55                                  _vU
	```
</details>

If the authentication was successful, the server will respond with an OK packet:
```
0000 0002 4f4b	....OK
```

### Heartbeating
The client sends a heartbeat every minute by just sending `00`.

### Join request
Once a client tries to join the server, the server will send you a packet with a length of 16, which seems to be some sort of connection token for use with the connection server.

## Connection server
When connecting to the connection server, the first packet sent will be the data segment of a join request from the communication server.

Once this is sent, the server will start proxying the client's packets to you and your packets to the client. You can then just directly talk to the minecraft client as if there was no proxy.
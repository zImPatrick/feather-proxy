# feather-proxy
feather-proxy is a standalone application that can be used to proxy traffic through [Feather Client](https://feathermc.com/)'s many proxies across the world. It simply acts like the official [Feather Client's "Player Server"](https://news.feathermc.com/host-a-free-minecraft-server/) feature.

## Usage
Currently, you have to get a server token yourself using the official Feather Client launcher.
0. Install Feather Client and sign in
1. Go to the "Server" section of the client (on the left)
2. Create a server with your desired proxy hostname, no other settings matter
3. Close the launcher and go to `%appdata%\.feather\player-server` and open `player-servers.json`  
It will look something like this:
```
[{
	"id": "02d658c0-233b-40f4-8ce1-d17138f0bb58",
	"template": "Vanilla",
	"version": "1.21.3",
	"slots": 20,
	"baseDirectory": "player-server/servers/02d658c0-233b-40f4-8ce1-d17138f0bb58",
	"serverToken": "<serverToken>",
	"ram": 2048,
	"launchCommand": "-Xms{RAM}M -Xmx{RAM}M"
}]
```

4. Copy the `serverToken` value and paste it into the `config.json` in this project. The config will automatically be generated once you run the program.
5. Run the program. It also has a few command line arguments, you can inspect these by appending `--help`. By default, the program will proxy to `127.0.0.1:25565`.

## Documentation
This repository also contains some documentation about Feather Client's api, if you're interested. It's in the [docs](./docs/) directory.
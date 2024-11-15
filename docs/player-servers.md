# Player servers
Player servers are a feather of Feather, with which you can host a server on your own computer and have it proxied behind Feather's network of reverse proxies.

## `POST https://api.feathermc.com/v1/player-server/search`
By using this endpoint, you can find all player servers you (and probably others? I haven't tested that yet) have created.

```http
POST /v1/player-server/search HTTP/1.1
Host: api.feathermc.com
Authorization: <authToken>
Content-Type: application/json

{
  "idOrMe": "me",
  "orderBy": {
    "label": "ASC"
  }
}
```

Example response:
```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8

{
  "total": 1,
  "results": [
    {
      "id": "3872b35f-046b-422b-b76d-fba54dbae5d3",
      "label": "test server",
      "proxyHostname": "testingserver",
      "lastAlive": "1970-01-01T00:00:00.000Z",
      "user": {
        "id": "2cd14c15-aabb-42b1-a9c4-bbc569f210f9",
        "username": null,
        "kind": "player",
        "email": null,
        "mcID": "8cedd991-9ffa-466d-8b05-bbfb49c5505d",
        "mcUsername": "notactuallyme"
      }
    }
  ]
}
```

## `POST https://api.feathermc.com/v1/player-server/create-server`
This endpoint is used to create a server on the backend. You'll get a server token back, which, as far as I know, **can't be retrieved after this**. Feather Client normally stores this in `%appdata%\.feather\player-servers\player-servers.json`. These tokens do not expire.

Example request:
```http
POST /v1/player-server/create-server HTTP/1.1
Host: api.feathermc.com
Content-Type: application/json
Authorization: <authToken>

{
  "label": "test server",
  "proxyHostname": "testingserver"
}
```

Example response:
```http
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8

{
  "id": "8b21579c-045d-449b-8a2f-51223abbc3ad",
  "proxyHostname": "testingserver",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiIyY2QxNGMxNS1hYWJiLTQyYjEtYTljNC1iYmM1NjlmMjEwZjkiLCJraW5kIjoicGxheWVyIiwidG9rZW5LaW5kIjoicGxheWVyX3NlcnZlciIsInRva2VuRGV0YWlscyI6eyJ0b2tlbklEIjoiMmM4NTQyMzAtMjM4NC00MzdiLTkyYjgtN2NlZjAxZmFjYzZmIiwic2VydmVySUQiOiIzODcyYjM1Zi0wNDZiLTQyMmItYjc2ZC1mYmE1NGRiYWU1ZDMifSwiaWF0IjoxNzIxMTY2MTAzfQ.i2YfO6dL39xD9gORojrD6KqVJ-lZW9lz7UEiETxb_vU"
}
```


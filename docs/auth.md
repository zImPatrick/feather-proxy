# Authentication with Feather Client
The official Feather launcher lets you log in with your official microsoft account - however, your minecraft auth token doesn't get sent to them in any way. Instead, they mimic how joining a server with the official minecraft client works:

tl;dr:
1. `GET https://api.feathermc.com/v1/minecraft/server-id` - Gets a server id which the launcher sends to Mojang
2. `POST https://sessionserver.mojang.com/session/minecraft/join`
3. `GET https://api.feathermc.com/v1/minecraft/has-joined/versetzt?token=<server-id>`
4. `GET https://api.feathermc.com/v1/game/auth-token`

This authentication flow gets you a token that can be used to communicate with all of Feather's services.

## 1. `GET https://api.feathermc.com/v1/minecraft/server-id`
No authentication or special headers needed. It just returns a random server id in the format of a JSON string.

Example response (without unnecessary headers):
```http
HTTP/1.1 201 Created
Content-Type: application/json

"d8f8c8395179b336ab282f13324171d299546175"
```

## 2. `POST https://sessionserver.mojang.com/session/minecraft/join`
The launcher then sends a join request to mojang with the server id we acquired from the previous request.

```http
POST /session/minecraft/join HTTP/1.1
Host: sessionserver.mojang.com
Content-Type: application/json
User-Agent: FeatherMC/Feather Client Launcher/1.6.1 (hello@feathermc.com)

{
  "accessToken": "<mojang access token>",
  "selectedProfile": "<profile uuid>",
  "serverId": "<server id we just got>"
}
```

Mojang just returns a `204 No Content` response. [The normal flow for joining a server is more thoroughly documented on wiki.vg](https://wiki.vg/Protocol_Encryption#Authentication). It's very similar, except that we get the server id directly from Feather's api.

## 3. `GET https://api.feathermc.com/v1/minecraft/has-joined/versetzt?token=<server-id>`
This is where Feather checks if you've sent a join request to Mojang.

Example of a successful response:
```http
HTTP/1.1 200 OK
authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiIyY2QxNGMxNS1hYWJiLTQyYjEtYTljNC1iYmM1NjlmMjEwZjkiLCJraW5kIjoicGxheWVyIiwibWNJRCI6IjhjZWRkOTkxLTlmZmEtNDY2ZC04YjA1LWJiZmI0OWM1NTA1ZCIsInRva2VuS2luZCI6InVzZXIiLCJpYXQiOjE3NjE2ODA5MzgsImV4cCI6MTc2MTkxMjMzNH0.CIFCY9j3JE8vKFDyE-LpyW_nr80t5oHS134E9-_vms

{
  "id": "2cd14c15-aabb-42b1-a9c4-bbc569f210f9", // Feather user id
  "username": null,
  "email": null,
  "kind": "player",
  "mcID": "8cedd991-9ffa-466d-8b05-bbfb49c5505d", // Minecraft uuid
  "mcUsername": "notactuallyme", // Minecraft username
  "firstName": null,
  "lastName": null,
  "registeredAt": "2024-15-11T11:16:36.444Z", // First feather registration
  "activatedAt": "2024-15-11T11:16:36.444Z", // When you first activated feather
  "country": "AT" // What country you're from
}
```

It also returns an `authorization` header (with a JWT with `userID`, `mcID`, `tokenKind`, but this is irrelevant for our usecase) which we will use in the next step to get a complete auth token.

## 4. `GET https://api.feathermc.com/v1/game/auth-token`
This is where we get a complete auth token. We pass in the authorization header we acquired in the previous step.

Example request:
```http
GET /v1/game/auth-token HTTP/1.1
Host: api.feathermc.com
Authorization: <authorization header from response of step 3>
```

Example response:
```json
"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjJjZDE0YzE1LWFhYmItNDJiMS1hOWM0LWJiYzU2OWYyMTBmOSIsInVzZXJuYW1lIjpudWxsLCJlbWFpbCI6bnVsbCwiZmlyc3ROYW1lIjpudWxsLCJsYXN0TmFtZSI6bnVsbCwicmVnaXN0ZXJlZEF0IjoiMjAyNC0xNS0xMVQxMToxNjozNi40NDRaIiwiYWN0aXZhdGVkQXQiOiIyMDI0LTE1LTExVDExOjE2OjM2LjQ0NFoiLCJraW5kIjoicGxheWVyIiwibWNJRCI6IjhjZWRkOTkxLTlmZmEtNDY2ZC04YjA1LWJiZmI0OWM1NTA1ZCIsIm1jVXNlcm5hbWUiOiJub3RhY3R1YWxseW1lIiwicmFua3MiOltdLCJwZXJtaXNzaW9ucyI6eyJwbGF5ZXJTZXJ2ZXJMaW1pdCI6M30sImlhdCI6MTc2MTY4MDkzOCwiZXhwIjoxNzYxOTEyMzM0fQ.ZmkR7apJaiHfmnu2GuFkRoC2W_KQ3xJIXVdNygf4GqU"
```
This is our authentication token that we include in an `Authorization` header on every request.

This JWT, when decoded, looks like this:
```json
{
  "id": "2cd14c15-aabb-42b1-a9c4-bbc569f210f9",
  "username": null,
  "email": null,
  "firstName": null,
  "lastName": null,
  "registeredAt": "2024-15-11T11:16:36.444Z",
  "activatedAt": "2024-15-11T11:16:36.444Z",
  "kind": "player",
  "mcID": "8cedd991-9ffa-466d-8b05-bbfb49c5505d",
  "mcUsername": "notactuallyme",
  "ranks": [],
  "permissions": {
    "playerServerLimit": 3
  },
  "iat": 1761680938,
  "exp": 1761912334
}
```
# Features

# Planned features
## Initial handshake
The inital handshake establishes key information between the client and server.
#### Server-to-Client
- **Version:** To ensure compatibility
- **Server type:** Confirms that its a game server
- **Session token:** Allows reconnecting clients to the servers saved state

#### Client-to-Server
- **Version:** Client version for compatibiliy check
- **Authentication token:** For security
## Gameloop
- UDP
## Player movement
- Update player coordinates, speed and direction

## Shooting
- Receive projectile information. Point of origin, direction and calculate if something is hit

## Chat
- Ingame chat for players to communicat
- TCP
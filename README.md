# Features
## **Packet Structure**

The server uses a custom binary protocol for communication. Each packet is structured as follows:

```
[Message Type: 1 byte][Payload Size: 4 bytes (big-endian)][Payload: variable length]
```

### **Header**
1. **Message Type (1 Byte)**:
   - Represents the type of message (e.g., `Handshake`, `Ping`, `Pong`, `Move`, etc.).
   - Encoded as an unsigned 8-bit integer (`uint8`).

2. **Payload Size (4 Bytes)**:
   - Indicates the size of the payload in bytes.
   - Encoded as an unsigned 32-bit integer (`uint32`) in **big-endian** format.

### **Payload**
- Contains the actual message data.
- Encoded as a UTF-8 string with a length matching the size specified in the header.

### **Supported Message Types**
| Message Type | Value            | Description            |
|--------------|------------------|------------------------|
| Handshake    | `0`              | Initial handshake      |
| Ping         | `1`              | Ping message           |
| Pong         | `2`              | Response to ping       |
| Move         | `3`              | Move command           |
| Shoot        | `4`              | Shoot command          |
| Chat         | `5`              | Chat message           |






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
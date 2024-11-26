# Features
## **Packet Structure**

The server uses a custom binary protocol for communication. Each packet is structured as follows:

```
[Message Type: 1 byte]
[Payload Size: 4 bytes]
[Serial: 1 byte]
[Error Flag: 1 byte]
[Timestamp: 8 bytes]
[Payload: variable length]
```

### **Header**
1. **Message Type (1 Byte)**:
   - Represents the type of message being sent or received.
   - Examples of message types:
     - `Handshake`: Initial client-server handshake.
     - `Ping`/`Pong`: Latency checking.
     - `Move`: Player movement data.
     - `Chat`: Text communication.
   - Encoded as an unsigned 8-bit integer (`uint8`).

2. **Payload Size (4 Bytes)**:
   - Indicates the size of the payload in bytes (excluding the header).
   - Encoded as an unsigned 32-bit integer (`uint32`) in **little-endian** format.

3. **Serial (1 Byte)**:
   - A unique identifier for the packet to enable ordering and detection of missing packets.
   - Encoded as an unsigned 8-bit integer (`uint8`), wrapping around at 255.

4. **Error Flag (1 Byte)**:
   - Indicates whether the message signals an error or is successful.
   - Values:
     - `0`: Success.
     - `1`: Error.
   - Encoded as an unsigned 8-bit integer (`uint8`).

5. **Timestamp (8 Bytes)**:
   - A UNIX timestamp (in milliseconds) indicating when the packet was created.
   - Encoded as a signed 64-bit integer (`int64`) in **little-endian** format.

#### **Payload**
- The payload contains the actual message data (e.g., player actions, chat messages).
- The size of the payload is determined by the `Payload Size` field in the header.

---

### **Packet Example**

Here’s an example of a packet in hexadecimal form:

```
01 00 00 00 0A 02 00 00 00 00 00 01 7E D5 48 61 6E 64 73 68 61 6B 65
```

#### Breakdown:
- `01`: Message Type (`Handshake`).
- `00 00 00 0A`: Payload Size (10 bytes).
- `02`: Serial Number (`2`).
- `00`: Error Flag (`0`, success).
- `00 00 00 00 00 01 7E D5`: Timestamp (`1690403200000`, corresponds to a UNIX time in milliseconds).
- `48 61 6E 64 73 68 61 6B 65`: Payload (`"Handshake"` in ASCII).

---

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
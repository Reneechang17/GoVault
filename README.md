# GoVault -- A Distributed File System

## Project Introduction
- Built a decentralized, content-addressable distributed file storage system in Go, featuring peer-to-peer communication, secure file storage, data broadcasting, file versioning, and encryption, designed to handle large-scale datasets across multiple nodes efficiently.

## Key Feature
### Decentralized, Peer-to-Peer Communication
- P2P Architecture: The system is designed with a decentralized peer-to-peer communication model, allowing nodes to communicate directly with each other without the need for a central server.
- TCP Transport: Utilized TCP connections for communication between peers.
- Peer Management: Added peer discovery and connection through a bootstrap mechanism and managed the lifecycle of connections across the network.

### Content-Addressable Storage
- File Storage with Hashing: Files are stored using content-addressable storage, where each file is assigned a unique key (hash) based on its content.
- File Write, Read, and Delete: Implemented core file operations like writing, reading, and deletion using the hashed file name for efficient lookup and management.

### Data Broadcasting and Synchronization
- Multi-Peer Broadcasting: Implemented data broadcasting across all connected peers, ensuring that file updates are synchronized across the network.
- gob Encoding: Used gob encoding and decoding for serialization and deserialization of data between nodes during transmission.

### File Versioning 
- Versioning: Introduced file versioning across the distributed network, ensured that each node has a consistent view of the data and can manage different versions of the same file.

### Security and Encryption
- Encrypted Communication: Implemented encryption to ensure secure file transfers between peers, protecting data integrity and privacy during transmission.
- Secure Storage: Planned for securely storing and managing files, with future scalability to support encrypted files on disk for enhanced security.

### Scalability for Large-Scale File Handling
- Scalable Design: The system is designed to handle large datasets efficiently, allowing for horizontal scaling as more nodes join the network.
- Efficient Resource Management: With the ability to store large files and manage resources across distributed nodes, the system is optimized for large-scale, decentralized file handling.

## Additional Features
- Modular Design: The system is built in a modular fashion with clear separation of concerns between components like file storage, peer communication, and message handling.
- Testing: Includes testing for core functionalities like file operations and peer connections, ensuring reliability and stability of the system.


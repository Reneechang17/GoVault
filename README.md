# GoVault -- A Distributed File Storage System

## 1. Project Description:

This project implements a **decentralized, peer-to-peer (P2P) distributed file storage system**. It is designed to allow nodes to communicate directly with each other, enabling secure file storage, retrieval, versioning, and encryption. The system uses a P2P architecture, ensuring scalability and fault tolerance by distributing the data across multiple nodes.

## 2. Key Features:

### 2.1 P2P Architecture:
The system utilizes a **decentralized P2P communication model** that allows nodes to communicate directly without relying on a central server. This ensures high availability and fault tolerance, as each node can serve as both a client and a server.

- **P2P** enables data sharing across distributed nodes.
- Each node can upload and download files, as well as share data with others.

### 2.2 TCP Transport:
Communication between nodes is achieved via **TCP transport**, providing reliable, ordered data transmission.

- **TCP connections** are used for establishing reliable communication channels between nodes.
- The system supports peer discovery and connection management using a bootstrap mechanism.
- The data is serialized using **GOB** encoding and decoded using custom decoders.

## 3. Encryption:

### 3.1 Decoder:
To handle data transmission, the system employs a custom **Decoder** interface. It ensures that data is decoded properly when received over the network.

- Two types of decoders are used: `GOBDecoder` for decoding regular messages, and `DefaultDecoder` for handling stream-based data (like file chunks).

### 3.2 Encryption/Decryption Process (AES/IV):
The system uses **AES** encryption in **CTR mode** to ensure secure file transfer between peers.

- A random **32-byte encryption key** is generated for each file, ensuring data confidentiality.
- **AES** is used to encrypt data, while an **Initialization Vector (IV)** is prepended to the encrypted data to ensure each encryption is unique.
- Files are encrypted and decrypted in streams, allowing for efficient handling of large files.

The encryption and decryption process involves:
- Generating an AES cipher using the key.
- Using the **CTR mode** for stream-based encryption/decryption.
- Managing the IV, which is essential for ensuring that identical data produces different encrypted outputs.

## 4. File Storage and Operations:

### 4.1 Path Hashing (SHA-1 & Path Segmentation):
The system uses **SHA-1 hashing** to generate unique file identifiers. This hashing ensures that file contents are uniquely represented and stored.

- **SHA-1**: Hashes the file key to create a unique identifier.
- **Path Segmentation**: The hash is split into smaller segments, which are used to construct the file's directory path. This approach avoids storing all files in a single directory and enhances scalability.

### 4.2 File Operations (Has/Write/Delete/Read):
- **Has**: Checks if the file exists in the local storage using the file's unique key.
- **Write**: Writes data to the local storage. Supports both plain and encrypted data.
- **Delete**: Deletes the file from the local storage.
- **Read**: Reads the file from the local storage and returns a readable stream of the file content.

These operations are implemented in the `Store` component, which manages the file storage and retrieval processes. The system ensures that file metadata and actual file content are correctly stored, retrieved, and deleted.

## 5. FileServer:
The **FileServer** is the core component that interacts with other nodes for storing and retrieving files. It:
- Manages file storage.
- Handles incoming and outgoing file requests.
- Performs encryption and decryption on the fly.
- Broadcasts file metadata to connected peers when a new file is uploaded.

### Key Methods:
- **Store**: Uploads files to the server and shares them with peers.
- **Get**: Retrieves a file, either from the local storage or by fetching it from a remote peer.

## 6. Module Design:
The system is designed to be modular, with clear separation of concerns between different components:

- **P2P Communication**: Handles peer connections and data exchange between nodes.
  - p2p/handshake.go, transport.go, tcp_transport.go, tcp_transport_test.go, message.go, encoding.go
- **File Storage**: Manages files, paths, and directory structures.
  - store.go, store_test.go
- **Encryption**: Ensures secure transmission and storage of files.
  - crypto.go, crypto_test.go
- **FileServer**: Orchestrates file-related operations like upload, download, and retrieval.
  - server.go

This modularity allows for flexibility and scalability, making it easy to add new features or replace components as needed.

## 7. Testing:
The system includes unit and integration tests to ensure reliability:

- **Path Transformation Test**: Ensures that file paths are correctly generated based on file keys.
- **File Operations Test**: Validates that files can be correctly written, read, and deleted.
- **Encryption/Decryption Test**: Verifies that files can be encrypted and decrypted without data loss or corruption.

These tests ensure that all components of the system function as expected, from file storage to encryption.

## 8. Future Improvements:
There are several areas where the system can be improved:

- **Security**: Transition from SHA-1 to a more secure hashing algorithm like **SHA-256** to avoid potential vulnerabilities.
- **Efficiency**: Improve the file transfer process by introducing parallel data transfers, especially for large files.
- **Scalability**: Implement **distributed hash tables (DHT)** for peer discovery and file location, which would enhance the scalability and fault tolerance of the network.
- **Error Handling**: Enhance error handling to better manage network failures and retries during file transfers.

## 9.Technical Skills and Patterns:
- Go, P2P architecture, TCP communication, encryption(AES, CTR Mode, SHA-1 Hashing), Serialization/Deserialization(GOB), concurrency and synchronization(Goroutines, Channels, Mutex)
# CA-DFS: Distributed, Peer-to-Peer Content Addressable Storage (CAS) System

CA-DFS is a distributed file system implemented in Go, offering a scalable and efficient solution for storing and retrieving data across a network of nodes. Designed as a peer-to-peer Content Addressable Storage (CAS) system, CA-DFS ensures high data integrity, deduplication, and redundancy across multiple servers.

## Features

- **Content Addressable Storage (CAS):** Ensures data integrity and deduplication by storing files based on their unique content hashes.
- **Peer-to-Peer Architecture:** Decentralized system allowing multiple nodes (servers) to store and retrieve files without a central authority.
- **Distributed System:** Horizontally scalable architecture that supports adding more nodes to increase storage capacity and fault tolerance.
- **Encryption:** Provides data security using AES encryption.
- **Custom Encoding:** Utilizes GOB encoding for efficient data serialization and transmission.
- **Flexible Transport Layer:** Built on TCP with potential for easy expansion to other protocols.

## System Architecture

CA-DFS is composed of several key components that work together to provide distributed, secure, and efficient file storage:

1. **FileServer:**  
   The core component responsible for managing file operations such as uploads, deletions, and retrievals. Each node runs a `FileServer` that interacts with peers to maintain data consistency across the network.

2. **Store:**  
   Handles local file storage on each server. Files are stored in a content-addressed format, and a customizable path transformation function organizes the file structure on disk.

3. **TCPTransport:**  
   Manages network communications between peers over TCP. This ensures reliable data transmission and synchronization across the distributed system.

## How It Works

### Uploading Files

When a file is uploaded to the main server, it is automatically distributed to other nodes in the network. The data is encoded using GOB encoding and encrypted with AES before being transferred. Each server in the network stores a piece of the encoded data, ensuring redundancy and fault tolerance.

### Deleting Files

When a file is deleted from the main server, the deletion request is propagated to all other nodes, removing the associated data from each node to maintain consistency across the network.

### Content Addressing

Files are stored based on their content hash, ensuring that the same file uploaded multiple times is stored only once. This mechanism also makes it easy to detect corrupted or tampered files.

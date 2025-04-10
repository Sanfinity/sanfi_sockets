# sanfi-sockets

**sanfi-sockets** is a collection of sample programs designed to help you learn and understand **WebSockets** in depth. This repository walks through basic to advanced concepts using hands-on code examples.

This is just a websocket exploring repository that contains some sample scripts that I've used to learn. 
---

## 🌐 What are WebSockets?

WebSockets are a communication protocol that provides **full-duplex**, **persistent** connections between a client (like a browser) and a server over a single TCP connection.

Unlike HTTP (which is request-response based), WebSockets allow **real-time communication** — meaning once the connection is established, data can be sent by either the client or server at any time without repeatedly reestablishing the connection.

---

## 🧠 Understanding WebSockets: From the Ground Up

### 📶 Network Layer Basics

In the OSI model, WebSockets operate over the **Application Layer** on top of the **Transport Layer (TCP)**.

Here’s a simplified view:

[Application Layer] → WebSocket [Transport Layer] → TCP [Network Layer] → IP [Data Link Layer] → Ethernet/Wi-Fi


- WebSockets use **TCP** for reliable delivery.
- The initial connection is made via **HTTP** (typically port 80 or 443).
- After a successful handshake, the protocol **upgrades** from HTTP to WebSocket.

---

## 🤔 Why Use WebSockets?

Here are some common reasons to use WebSockets:

| Feature              | HTTP                   | WebSocket                |
|----------------------|------------------------|--------------------------|
| Direction            | Client → Server only   | Full-duplex (2-way)      |
| Connection           | New request every time | Persistent               |
| Latency              | Higher                 | Very low                 |
| Use case             | Standard web browsing  | Real-time apps (chat, games, live data) |

### ✅ Use Cases

- **Chat applications**
- **Live sports scores**
- **Collaborative editing (e.g. Google Docs)**
- **Online multiplayer games**
- **Real-time notifications**

---

## 💻 A Simple WebSocket Example

Let’s look at a super basic WebSocket setup using **Node.js** and the **ws** library:

### Server (`server.js`)

```js
const WebSocket = require('ws');

const server = new WebSocket.Server({ port: 8080 });

server.on('connection', socket => {
  console.log('Client connected');

  socket.on('message', message => {
    console.log('Received:', message);
    socket.send(`Echo: ${message}`);
  });

  socket.on('close', () => console.log('Client disconnected'));
});
```

### Client (HTML + JS)

```js
<!DOCTYPE html>
<html>
<body>
  <input id="msg" placeholder="Type a message..." />
  <button onclick="send()">Send</button>
  <pre id="log"></pre>

  <script>
    const socket = new WebSocket('ws://localhost:8080');
    const log = document.getElementById('log');

    socket.onmessage = e => log.textContent += 'Server: ' + e.data + '\n';

    function send() {
      const msg = document.getElementById('msg').value;
      socket.send(msg);
      log.textContent += 'You: ' + msg + '\n';
    }
  </script>
</body>
</html>
```
---
## 🚀 Getting Started
Clone the repo and install dependencies:

```bash
git clone https://github.com/yourusername/sanfi-sockets.git
cd sanfi-sockets
```
Open the files in your editor and start experimenting!

---
## 🌍 Popular WebSocket Applications and Frameworks by Language

Here are some trending and widely-used WebSocket implementations and applications across different programming languages:

### 🟦 JavaScript / Node.js

- **Socket.IO** – Simplifies WebSocket usage with fallbacks and rooms.
- **ws** – A fast, simple, and lightweight WebSocket library for Node.js.
- **Primus** – A flexible framework with support for multiple transport layers.

🔗 Real-world use:  
**Slack**, **Trello**, and **Discord** use WebSocket-based architectures for real-time communication.

---

### 🟨 Python

- **FastAPI + WebSockets** – Modern framework with native WebSocket support.
- **Django Channels** – Adds WebSocket and async support to Django.
- **websockets** – A lightweight WebSocket library for building clients and servers.

🔗 Real-world use:  
**Anvil** and **Home Assistant** use WebSockets for real-time data and user interfaces.

---

### 🟥 Java

- **Spring WebSocket** – Integrated with Spring Boot, supports STOMP over WebSockets.
- **Netty** – Event-driven networking for real-time apps.
- **Atmosphere** – Supports WebSockets and other transport methods.

🔗 Real-world use:  
Many **banking dashboards** and **trading platforms** built on Java use Spring WebSockets.

---

### 🟪 C# / .NET

- **ASP.NET Core SignalR** – Simplifies adding real-time web functionality.
- **WebSocketSharp** – A .NET library supporting both server and client.

🔗 Real-world use:  
**Microsoft Teams** uses WebSockets heavily for messaging and presence updates.

---

### 🟩 Go (Golang)

- **Gorilla WebSocket** – The most popular and performant WebSocket library for Go.
- **nhooyr/websocket** – A minimal and idiomatic WebSocket library.

🔗 Real-world use:  
Used in **monitoring dashboards**, **real-time logs**, and **IoT backends**.

---

### 🟫 Rust

- **tokio-tungstenite** – Asynchronous WebSocket library for Tokio.
- **warp + WebSocket** – A modern web framework that integrates easily with WebSockets.

🔗 Real-world use:  
**Real-time multiplayer games**, **blockchain nodes**, and **sensor networks** are emerging use cases.

---

### 🐘 PHP

- **Ratchet** – A PHP library for real-time bi-directional applications.
- **Workerman** – High-performance asynchronous WebSocket server in PHP.

🔗 Real-world use:  
Used in **live chat widgets**, **collaborative editing**, and **push notifications** for PHP-based platforms.

---

## 📔 References:
- https://websockets.readthedocs.io/en/stable/intro/tutorial1.html
- https://www.youtube.com/watch?v=8ARodQ4Wlf4

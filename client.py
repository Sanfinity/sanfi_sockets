import asyncio
import websockets

async def communicate():
    async with websockets.connect("ws://localhost:8765") as websocket:
        while True:
            message = input("Enter message (or 'quit' to exit): ")
            if message.lower() == 'quit':
                break
            
            await websocket.send(message)
            print(f"Sent: {message}")
            
            response = await websocket.recv()
            print(f"Received: {response}")

asyncio.run(communicate())
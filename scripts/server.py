import asyncio
import websockets

async def echo(websocket):
    async for message in websocket:
        await websocket.send("Server says " + message)

async def main():
    async with websockets.serve(echo, "localhost", 8765):
        await asyncio.Future()  # run forever

asyncio.run(main())
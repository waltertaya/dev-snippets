# server.py
import asyncio
import websockets

connected_clients = set()

async def handle_client(websocket):
    # Register the client
    connected_clients.add(websocket)
    print(websocket)
    print(connected_clients)
    try:
        async for message in websocket:
            # Broadcast the message to all connected clients
            for client in connected_clients:
                if client != websocket:  # Avoid echoing back to the sender
                    await client.send(message)
    except websockets.ConnectionClosed:
        print("Client disconnected")
    finally:
        # Unregister the client
        connected_clients.remove(websocket)

async def main():
    print("WebSocket Server is running on ws://localhost:8765")
    async with websockets.serve(handle_client, "localhost", 8765):
        await asyncio.Future()  # Run forever

if __name__ == "__main__":
    asyncio.run(main())

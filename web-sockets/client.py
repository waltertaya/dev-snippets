# client.py
import asyncio
import websockets

async def chat_client():
    uri = "ws://localhost:8765"
    async with websockets.connect(uri) as websocket:
        print("Connected to the chat server.")
        print("Type your messages below:")
        
        async def send_messages():
            while True:
                message = input("You: ")
                await websocket.send(message)
        
        async def receive_messages():
            while True:
                message = await websocket.recv()
                print(message)
                print(f"\rFriend: {message}\nYou: ", end="")  # Update terminal without breaking input
        
        # Run send and receive tasks concurrently
        await asyncio.gather(send_messages(), receive_messages())

if __name__ == "__main__":
    asyncio.run(chat_client())

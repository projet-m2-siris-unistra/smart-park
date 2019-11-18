import asyncio
from nats.aio.client import Client as NATS
import sys

async def run():
    nc = NATS()
    await nc.connect()

    response = await nc.request(sys.argv[1], bytes(sys.argv[2], 'utf-8'), timeout=1)
    print(response.data)

    await nc.close()


loop = asyncio.get_event_loop()
loop.run_until_complete(run())
loop.close()

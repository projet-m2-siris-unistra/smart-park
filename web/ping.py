import asyncio
import sys

from nats.aio.client import Client as NATS


async def run():
    nc = NATS()
    await nc.connect()

    payload = b""
    if len(sys.argv) == 3:
        payload = bytes(sys.argv[2], "utf-8")

    response = await nc.request(sys.argv[1], payload, timeout=1)
    print(response.data)

    await nc.close()


loop = asyncio.get_event_loop()
loop.run_until_complete(run())
loop.close()

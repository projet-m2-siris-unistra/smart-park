import asyncio
from nats.aio.client import Client as NATS


async def run():
    nc = NATS()
    await nc.connect()

    response = await nc.request("ping", b"", timeout=1)
    print(response.data)

    await nc.close()


loop = asyncio.get_event_loop()
loop.run_until_complete(run())
loop.close()

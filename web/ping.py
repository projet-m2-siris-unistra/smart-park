import json
import asyncio
import sys

from nats.aio.client import Client as NATS
from nats.aio.errors import ErrTimeout

async def run():
    nc = NATS()
    await nc.connect()

    payload = b""
    if len(sys.argv) == 3:
        payload = bytes(sys.argv[2], "utf-8")
    try:
        response = await nc.request(sys.argv[1], payload, timeout=1)
        data = json.loads(response.data)
        print(json.dumps(data, indent=2))
    except ErrTimeout:
        print("Timeout atteint : stop process")

    await nc.close()


loop = asyncio.get_event_loop()
loop.run_until_complete(run())
loop.close()

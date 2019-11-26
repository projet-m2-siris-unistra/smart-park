import logging

from nats.aio.client import Client as NATS
import json

logger = logging.getLogger(__name__)
nc = NATS()


async def error_cb(err):
    logger.error("error while connecting to NATS: " + str(err))


async def init(app, loop):
    await nc.connect(app.config.NATS_URL, loop=loop, error_cb=error_cb)
    app.nc = nc


async def close(app, loop):
    await nc.close()


async def ping():
    response = await nc.request("ping", b"", timeout=1)
    return response.data.decode("utf-8")

async def getTenant(tenant_id):
    request = json.dumps({'tenant_id' : tenant_id})
    response = await nc.request("tenants.get", bytes(request, "utf-8"), timeout=1)
    return response.data.decode("utf-8")

def setup(app):
    app.register_listener(init, "before_server_start")
    app.register_listener(close, "after_server_stop")

from nats.aio.client import Client as NATS

nc = NATS()


async def init(app, loop):
    await nc.connect(app.config.NATS_URL, loop=loop)
    app.nc = nc


async def close(app, loop):
    await nc.close()


async def ping():
    response = await nc.request("ping", b"", timeout=1)
    return response.data.decode("utf-8")


def setup(app):
    app.register_listener(init, "before_server_start")
    app.register_listener(close, "after_server_stop")

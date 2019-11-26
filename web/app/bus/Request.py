from nats.aio.client import Client as NATS
from nats.aio.errors import ErrTimeout

import json

from app.bus import nc


# Request tenant infos
async def getTenant(tenant_id):
    request = json.dumps({'tenant_id' : tenant_id})
    response = await nc.request("tenants.get", bytes(request, "utf-8"), timeout=1)
    return response.data.decode("utf-8")

# Request zone infos
async def getZone(zone_id):
    request = json.dumps({'zone_id' : zone_id})
    response = await nc.request("zones.get", bytes(request, "utf-8"), timeout=1)
    return response.data.decode("utf-8")

# This will return a list of zones from a tenant
async def getZones(tenant_id):
    response = await nc.request("zones.list", b"{}", timeout=1)
    return response.data.decode("utf-8")


# Request spot infos
async def getSpot(spot_id):
    request = json.dumps({'zone_id' : spot_id})
    print("Request to DB:" + request)
    try: 
        response = await nc.request("spot", bytes(request, "utf-8"), timeout=1)
    except ErrTimeout:
        print("Request timed out")
    return format(response.data.decode())

"""
Further requests will get infos on topics like
    zone.devices
"""

# Request spot infos
async def getDevice(device_id):
    request = json.dumps({'tenant_id' : device_id})
    response = await nc.request("devices.get", bytes(request, "utf-8"), timeout=1)
    return response.data.decode("utf-8")

async def getDevicesList():
    response = await nc.request("devices.list", b"{}", timeout=1)
    return response.data.decode("utf-8")
from nats.aio.client import Client as NATS
from nats.aio.errors import ErrTimeout

import json

from app.bus import nc


# Request tenant infos
async def getTenant(tenant_id):
    request = json.dumps({'tenant_id' : int(tenant_id)})
    response = await nc.request("tenants.get", bytes(request, "utf-8"), timeout=1)
    return response.data.decode("utf-8")


# Request zone infos
async def getZone(zone_id):
    request = json.dumps({'zone_id' : int(zone_id)})
    response = await nc.request("zones.get", bytes(request, "utf-8"), timeout=1)
    return response.data.decode("utf-8")


# This will return a list of zones from a tenant
async def getZones(tenant_id):
    request = json.dumps({'tenant_id' : int(tenant_id)})
    try:
        response = await nc.request("zones.list", bytes(request, "utf-8"), timeout=1)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> getZones -> timeout reached")
    return response.data.decode("utf-8")


# Returns all the spots associated to zone_id
async def getSpots(zone_id):
    request = json.dumps({'zone_id' : int(zone_id)})
    response = await nc.request("places.list", bytes(request, "utf-8"), timeout=1)
    return response.data.decode("utf-8")


# Request spot infos
async def getSpot(spot_id):
    request = json.dumps({'place_id' : int(spot_id)})
    response = await nc.request("places.get", bytes(request, "utf-8"), timeout=1)
    return response.data.decode("utf-8")


# Request spot infos
async def getDevice(device_id):
    request = json.dumps({'device_id' : int(device_id)})
    response = await nc.request("devices.get", bytes(request, "utf-8"), timeout=1)
    return response.data.decode("utf-8")


async def getDevicesList():
    response = await nc.request("devices.list", b"{}", timeout=1)
    return response.data.decode("utf-8")
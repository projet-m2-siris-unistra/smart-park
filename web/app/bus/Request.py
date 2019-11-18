from nats.aio.client import Client as NATS
from nats.aio.errors import ErrTimeout

import json

from app.bus import nc


# Request tenant infos + zone list
async def getTenant(tenant_id):
    request = json.dumps({'tenant_id' : tenant_id})
    print("Request to DB:" + request)
    try: 
        response = await nc.request(
            "tenant", bytes(request, "utf-8"), timeout=1)
    except ErrTimeout:
        print("Request timed out")
    return format(response.data.decode())
    

# Request zone infos + spot list
async def getZone(zone_id):
    request = json.dumps({'zone_id' : zone_id})
    print("Request to DB:" + request)
    try: 
        response = await nc.request(
            "zone", bytes(request, "utf-8"), timeout=1)
    except ErrTimeout:
        print("Request timed out")
    return format(response.data.decode())


# Request spot infos + devices list
async def getZone(spot_id):
    request = json.dumps({'zone_id' : spot_id})
    print("Request to DB:" + request)
    try: 
        response = await nc.request(
            "spot", bytes(request, "utf-8"), timeout=1)
    except ErrTimeout:
        print("Request timed out")
    return format(response.data.decode())


# Request spot infos + devices list
async def getDevice(device_id):
    request = json.dumps({'device_id' : device_id})
    print("Request to DB:" + request)
    try: 
        response = await nc.request(
            "device", bytes(request, "utf-8"), timeout=1)
    except ErrTimeout:
        print("Request timed out")
    return format(response.data.decode())
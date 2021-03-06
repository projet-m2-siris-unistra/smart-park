from nats.aio.client import Client as NATS
from nats.aio.errors import ErrTimeout

import json

from app.bus import nc

# Return values for request exception handling
REQ_ERROR = -1
REQ_OK = 0

# Tolerated delay (in seconds) when making request to the bus
globalTimeout = 1


"""
NOTE
Warning: some attributs do not have the same naming in the front-end
as in the backend (e.g. coordinates & geo)
"""



# Request tenant infos
async def getTenant(tenant_id):
    request = json.dumps({'tenant_id' : int(tenant_id)})
    try:
        response = await nc.request("tenants.get", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> getTenant -> timeout reached")
        return REQ_ERROR
    
    return response.data.decode("utf-8")


# Request zone infos
async def getZone(zone_id):
    request = json.dumps({
        'zone_id' : int(zone_id),
        'with_places' : True
    })
    try:
        response = await nc.request("zones.get", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> getZone -> timeout reached")
        return REQ_ERROR

    return response.data.decode("utf-8")


# This will return a list of zones from a tenant
async def getZones(tenant_id, page, pagesize, with_places=False):
    # calculating offset
    offset = pagesize * (page-1)

    request = json.dumps({
        'tenant_id' : int(tenant_id),
        'limit' : pagesize,
        'offset' : offset,
        'with_places' : with_places
    })

    print("Request : ", request)

    try:
        response = await nc.request("zones.list", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> getZones -> timeout reached")
        return REQ_ERROR

    return response.data.decode("utf-8")


# Add a zone to database
async def createZone(tenant_id, name, type, color, polygon):
    request = json.dumps({
        'tenant_id' : int(tenant_id),
        'name' : name,
        'type' : type,
        'color' : color,
        'geo' : polygon
    })
    print("createZone request = ", request)
    
    try:
        response = await nc.request("zones.new", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> createZones -> timeout reached")
        return REQ_ERROR

    return REQ_OK


# Update a zone
async def updateZone(zone_id, tenant_id, name, type, color, polygon):
    request = json.dumps({
        'zone_id' : int(zone_id),
        'tenant_id' : int(tenant_id),
        'name' : name,
        'type' : type,
        'color' : color,
        'polygon' : polygon
    })
    print("updateZone request = ", request)

    try:
        response = await nc.request("zones.update", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> updateZone -> timeout reached")
        return REQ_ERROR
    
    return REQ_OK


async def deleteZone(zone_id):
    request = json.dumps({
        'zone_id' : int(zone_id)
    })
    print("deleteZone request = ", request)
    
    try:
        response = await nc.request("zones.delete", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> deleteZone -> timeout reached")
        return REQ_ERROR
    
    return REQ_OK


# Returns all the spots associated to zone_id
async def getSpots(zone_id, page, pagesize):
    offset = pagesize * (page-1)
    
    request = json.dumps({
        'zone_id' : int(zone_id),
        'limit' : pagesize,
        'offset' : offset
    })

    try:
        response = await nc.request("places.list", bytes(request, "utf-8"))
    except ErrTimeout:
        print("WARNING: bus/Request.py -> getSpots -> timeout reached")
        return REQ_ERROR
    
    return response.data.decode("utf-8")


# Request spot infos
async def getSpot(spot_id):
    request = json.dumps({'place_id' : int(spot_id)})
    try:
        response = await nc.request("places.get", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> getSpot -> timeout reached")
        return REQ_ERROR
    
    return response.data.decode("utf-8")


# Creatng a spot
async def createSpot(zone_id, device_id, type, coordinates):
    print("coordinates: ", coordinates)
    request = json.dumps({
        'zone_id' : int(zone_id),
        'type' : type,
        'device_id' : device_id,
        'geo' : str(coordinates)
    })
    print("createSpot request = ", request)
    
    try:
        response = await nc.request("places.new", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> createSpot -> timeout reached")
        return REQ_ERROR

    print("RESPONSE: ", response.data.decode("utf-8"))
    return REQ_OK


async def updateSpot(spot_id, device_id, type, coordinates):
    request = json.dumps({
        'place_id' : int(spot_id),
        'device_id' : int(device_id),
        'type' : type,
        'geo' : str(coordinates)
    })
    print("updateSpot request = ", request)

    try:
        response = await nc.request("places.update", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> updateSpot -> timeout reached")
        return REQ_ERROR
    print("response: ", response.data.decode("utf-8"))
    return REQ_OK


async def deleteSpot(spot_id):
    request = json.dumps({
        'place_id' : int(spot_id)
    })
    print("deleteSpot request = ", request)
    
    try:
        response = await nc.request("places.delete", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> deleteSpot -> timeout reached")
        return REQ_ERROR
    
    return REQ_OK


# Request spot infos
async def getDevice(device_id):
    request = json.dumps({'device_id' : int(device_id)})
    try:
        response = await nc.request("devices.get", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> getDevice -> timeout reached")
        return REQ_ERROR

    return response.data.decode("utf-8")


# Request all devices of a tenant
async def getDevices(tenant_id, page, pagesize):
    # calculating offset
    offset = pagesize * (page-1)

    request = json.dumps({
        'tenant_id' : int(tenant_id),
        'limit' : pagesize,
        'offset' : offset
    })

    print("Request : ", request)

    try:
        response = await nc.request("devices.list", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> getDevices -> timeout reached")
        return REQ_ERROR

    return response.data.decode("utf-8")


# Request all NOT ASSIGNED devices
async def getNotAssignedDevices(tenant_id, page, pagesize):
    # calculating offset
    offset = pagesize * (page-1)

    request = json.dumps({
        'tenant_id' : int(tenant_id),
        'limit' : pagesize,
        'offset' : offset
    })
    
    try:
        response = await nc.request("devices.get.notassigned", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        return REQ_ERROR

    return response.data.decode("utf-8")


# This variable contains the values that device.state can take
deviceStates = {
    'free' : 'free',
    'occupied' : 'occupied',
    'notassigned' : 'notassigned'
}

# Create a device from EUI
async def createDevice(tenant_id, eui, name):
    # Name is not used for now. Waiting for migration update
    
    # default values that will (hopefully) be updated from 
    # device communication or tenant configuration
    batteryDefault = 100
    stateDefault = deviceStates['free']

    request = json.dumps({
        'tenant_id' : tenant_id,
        'device_eui' : eui,
        'battery' : batteryDefault,
        'state' : stateDefault
    })

    print("INFO: request=", request)

    try:
        response = await nc.request("devices.new", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        return REQ_ERROR

    print("INFO: reponse = ", response.data.decode("utf-8"))
    print("INFO: device created with device_id=", 
        json.loads(response.data.decode("utf-8"))['device_id'])
    return REQ_OK


async def deleteDevice(device_id):
    request = json.dumps({
        'device_id' : int(device_id)
    })
    print("deleteDevice request = ", request)
    
    try:
        response = await nc.request("devices.delete", bytes(request, "utf-8"), timeout=globalTimeout)
    except ErrTimeout:
        print("WARNING: bus/Request.py -> deleteDevice -> timeout reached")
        return REQ_ERROR
    
    return REQ_OK
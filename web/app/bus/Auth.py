from nats.aio.client import Client as NATS
from nats.aio.errors import ErrTimeout

import json

from app.bus import nc
from app.bus.Request import REQ_ERROR, REQ_OK



async def authorization(config, url):
    request = json.dumps({
        "config" : config, 
        "redirect_uri" : url
    })
    
    res = await nc.request("auth.authorization", bytes(request, "utf-8"))
    
    if res == REQ_ERROR:
        raise ServerError("auth.authorization failed")
    
    return res.data.decode("utf-8")


async def exchange(config, code, url):
    request = json.dumps({
        "config" : config, 
        "code" : code, 
        "redirect_uri" : url
    })
    
    res = await nc.request("auth.exchange", bytes(request, "utf-8"))
    
    if res == REQ_ERROR:
        return REQ_ERROR
    
    return res.data.decode("utf-8")


async def validate(token):
    request = json.dumps({
        "token" : token
    })
    
    res = await nc.request("auth.validate", bytes(request, "utf-8"))
    
    if res == REQ_ERROR:
        return REQ_ERROR
    
    return res.data.decode("utf-8")

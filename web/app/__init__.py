import click
import asyncio
from jinja2 import Environment, PackageLoader, select_autoescape
from sanic import Sanic, response
from sanic.response import json
from sanic.handlers import ErrorHandler
from sanic.exceptions import ServerError 
from sanic_session import Session, InMemorySessionInterface

import json as js 

import app.accounts
import app.bus
import app.config
import app.zones
import app.spots
import app.devices

from app.templating import render
from app.parkings import TenantManagement
from app.parkings import ZoneManagement
from app.parkings import Tooling
from app.pagination import Pagination

app = Sanic(__name__)
session = Session(app, interface=InMemorySessionInterface())
app.register_listener(config.load, "before_server_start")
bus.setup(app) # app.nc ==> connected object

app.config['WTF_CSRF_SECRET_KEY'] = '*Dieu Quentin*'

app.static("/static", "./static")

app.blueprint(accounts.bp)
app.blueprint(zones.bp)
app.blueprint(spots.bp)
app.blueprint(devices.bp)



# Handling navigation
@app.route("/")
@app.route("/home")
async def home(request):
    rendered_template = await render(
        "base_template.html", 
        request, 
        knights="BIENVENUE SUR SMART PARK !"
    )
    return response.html(rendered_template)


@app.route("/dashboard")
async def dashboard(request):
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)

    await tenantInstance.setZones()
    if tenantInstance.zones is None:
        raise ServerError("No zones in DB", status_code=500)

    zonesJson = Tooling.jsonList(tenantInstance.zones)

    rendered_template = await render(
        "dashboard_template.html",
        request, 
        tenantInstance=tenantInstance
    )
    return response.html(rendered_template)


@app.route("/zones")
async def zones(request):
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)

    # Calculating and initialization of pagination
    pagination = Pagination(request)
    
    # Requesting zone from database
    res = await tenantInstance.setZones(
        page=pagination.page_number, 
        pagesize=pagination.page_size
    )
    if tenantInstance.zones is None:
        raise ServerError("No zones in DB", status_code=500)
    
    pagination.setElementsNumber(tenantInstance.zonesCount)

    zonesJson = Tooling.jsonList(tenantInstance.zones)

    rendered_template = await render(
        "tenant_zone_data_table.html", 
        request,
        tenantInstance=tenantInstance,
        zonesList=zonesJson,
        pagination = pagination
    )
    return response.html(rendered_template)


@app.route("/map")
async def map(request):
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)

    await tenantInstance.setZones()
    if tenantInstance.zones is None:
        raise ServerError("No zones in DB", status_code=500)

    for zone in tenantInstance.zones:
        await zone.setSpots()

    zonesJson = Tooling.jsonList(tenantInstance.zones)
    
    rendered_template = await render(
        "map_template.html",
        request,
        tenantInstance=tenantInstance,
        zoneList=zonesJson
    )
    return response.html(rendered_template)


@app.route("/configuration")
async def configuration(request):
    rendered_template = await render(
        "base_template.html", 
        request, 
        knights="En cours de construction..."
    )
    return response.html(rendered_template)


@app.route("/statistics")
async def statistics(request):
    rendered_template = await render(
        "base_template.html", 
        request, 
        knights="En cours de construction..."
    )
    return response.html(rendered_template)


@app.route("/ping")
async def ping(request):
    ret = await bus.ping()
    return response.json({"data": ret})


@click.command()
@config.run_params
def run():
    server = app.create_server(host=config.env("HOST"), port=config.env("PORT"), debug=config.env("DEBUG"), return_asyncio_server=True)
    # NATS with TLS does not work properly with uvloop for some reason
    # Sanic changes the default event loop by default if uvloop is installed
    asyncio.set_event_loop_policy(None)
    loop = asyncio.get_event_loop()
    task = asyncio.ensure_future(server)
    loop.run_forever()


# Errors & Exceptions handling
@app.exception(ServerError)
async def serverErrorHandler(request, exception):
    print(exception)
    rendered_template = await render(
        "server_error_template.html", 
        request,
        errCode = exception,
        knights = "Veuillez nous excuser, une erreur interne vient de se produire...",
    )
    return response.html(rendered_template)

import click
from jinja2 import Environment, PackageLoader, select_autoescape
from sanic import Sanic, response
from sanic.response import json
from sanic.handlers import ErrorHandler
from sanic.exceptions import ServerError 
from sanic_session import Session, InMemorySessionInterface

import json as js 
from math import ceil

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

    # number of total elements available in DB
    limit = 20
    offset = 1 
    
    print("req=", request.raw_args)
    
    # number of elements per page
    if "pagesize" in request.raw_args:
        limit = int(request.raw_args['pagesize'])
        if limit not in [10, 20, 30, 40, 50]:
            print("WARNING: pagesize is not valid")
            raise ServerError("page size not valid")
    
    # current page
    if "page" in request.raw_args:
        offset = int(request.raw_args['page'])
        if offset < 1: #or offset > count: ==> CAN'T TEST IT HERE !
            print("WARNING: page number is not valid")
            raise ServerError("page number not valid")

    res = await tenantInstance.setZones(page=offset, pagesize=limit)
    if tenantInstance.zones is None:
        raise ServerError("No zones in DB", status_code=500)
    
    # Number of pages available (NOTE: only the current page is loaded)
    count = tenantInstance.zonesCount
    pages = ceil(count/limit)

    zonesJson = Tooling.jsonList(tenantInstance.zones)

    rendered_template = await render(
        "tenant_zone_data_table.html", 
        request,
        tenantInstance=tenantInstance,
        zonesList=zonesJson,
        paginationLimit=limit,
        paginationOffset=offset,
        paginationElements=count,
        paginationPages=pages
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
    app.run(host=config.env("HOST"), port=config.env("PORT"), debug=config.env("DEBUG"))


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

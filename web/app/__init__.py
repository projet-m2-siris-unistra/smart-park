import click
from jinja2 import Environment, PackageLoader, select_autoescape
from sanic import Sanic, response
from sanic.response import json

import json as js

import app.accounts
import app.bus
import app.config
import app.zones
import app.spots

from app.templating import render
from app.parkings import TenantManagement
from app.parkings import ZoneManagement
from app.parkings import Tooling

app = Sanic(__name__)
app.register_listener(config.load, "before_server_start")
bus.setup(app) # app.nc ==> connected object

app.static("/static", "./static")

app.blueprint(accounts.bp)
app.blueprint(zones.bp)
app.blueprint(spots.bp)


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
    zonesJson = Tooling.jsonList(tenantInstance.zones)

    rendered_template = await render(
        "dashboard_template.html",
        request, 
        zoneList = tenantInstance.zones,
        totalSpots = tenantInstance.getTotalSpots(),
        takenSpots = tenantInstance.getTakenSpots()
    )
    return response.html(rendered_template)


@app.route("/zones")
async def zones(request):
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)

    await tenantInstance.setZones()

    rendered_template = await render(
        "tenant_zone_data_table.html", 
        request,
        zoneList=tenantInstance.zones,
        tenantName=tenantInstance.name
    )
    return response.html(rendered_template)


@app.route("/map")
async def map(request):
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)

    await tenantInstance.setZones()
    zonesJson = Tooling.jsonList(tenantInstance.zones)

    rendered_template = await render(
        "map_template.html",
        request,
        tenantName=tenantInstance.name,
        tenantCoor=tenantInstance.coordinates,
        zoneJsonList=zonesJson
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


@app.route("/devices")
async def devices(request):
    # Going through hierarchy for testing
    devicesList = await bus.Request.getDevicesList()
    devicesListJson = js.loads(devicesList)
    rendered_template = await render(
        "devices_template.html", 
        request,
        devices = devicesListJson
    )
    return response.html(rendered_template)


@app.route("/ping")
async def ping(request):
    ret = await bus.ping()
    return response.json({"data": ret})


# Testing
@app.route("/getTenant")
async def getTenant(request):
    ret = await bus.getTenant(1)
    data = js.loads(ret)
    return response.json(data["geo"])


@click.command()
@config.run_params
def run():
    app.run(host=config.env("HOST"), port=config.env("PORT"), debug=config.env("DEBUG"))

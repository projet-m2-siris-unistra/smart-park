import click
from jinja2 import Environment, PackageLoader, select_autoescape
from sanic import Sanic, response
from sanic.response import json

import app.accounts
import app.bus
import app.config
import app.zones

from app.templating import render
from app.parkings import TenantManagement

app = Sanic(__name__)
app.register_listener(config.load, "before_server_start")
bus.setup(app) # app.nc ==> connected object

app.static("/static", "./static")

app.blueprint(accounts.bp)
app.blueprint(zones.bp)


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
    # request from ID
    tenantInstance = TenantManagement(123)
    rendered_template = await render(
        "dashboard_template.html",
        request, 
        zoneList = tenantInstance.getZones(),
        totalSpots = tenantInstance.getTotalSpots(),
        takenSpots = tenantInstance.getTakenSpots()
    )
    return response.html(rendered_template)


@app.route("/zones")
async def zones(request):
    rendered_template = await render("zones_template.html", request)
    return response.html(rendered_template)


@app.route("/map")
async def map(request):
    rendered_template = await render("map_template.html", request)
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

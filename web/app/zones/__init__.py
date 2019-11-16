import json as js

from sanic import Blueprint, response
from sanic.response import json

from app.templating import render

from app.parkings import ZoneManagement
from app.parkings import TenantManagement

bp = Blueprint("zones", url_prefix='/parking')

# Handling Parkings zones
@bp.route('/create_zone')
async def create_zone(request):
    rendered_template = await render("zone_creation.html", request)
    return response.html(rendered_template)

@bp.route('/zone-creation-check', methods=['POST'])
async def zone_creation_check(request):
    # Checking args
    # Adding args to database
    # Linking to zone management
    return response.json(request.form)

@bp.route('/<zone>')
@bp.route('/<zone>/overview')
async def view(request, zone):
    tenant = TenantManagement(123)
    zoneInstance = ZoneManagement(zone)
    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_view='true', 
        zoneName=zone,
        tenantName=tenant.name,
        zoneTaken=zoneInstance.getNbTakenSpots(),
        zoneTotal=zoneInstance.getNbTotalSpots()
    )
    return response.html(rendered_template)

@bp.route('/<zone>/statistics')
async def stats(request, zone):
    tenant = TenantManagement(123)
    zoneInstance = ZoneManagement(zone)
    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_stats='true', 
        zoneName=zone,
        tenantName=tenant.name,
        statistics=zoneInstance.getAllStats()
    )
    return response.html(rendered_template)

@bp.route('/<zone>/maintenance')
async def maintenance(request, zone):
    tenant = TenantManagement(123)
    zoneInstance = ZoneManagement(zone)
    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_maintenance='true',
        zoneName=zone,
        tenantName=tenant.name,    
        spotList=zoneInstance.getSpotList()
    )
    return response.html(rendered_template)

@bp.route('/<zone>/configuration')
async def config(request, zone):
    tenant = TenantManagement(123)
    zoneInstance = ZoneManagement(zone)
    spotsJson = jsonList(zoneInstance.getSpotList())
    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_config='true',
        zoneName=zone,
        tenantName=tenant.name,    
        zonePolygon=zoneInstance.getPolygon(),
        spotList=spotsJson
    )
    return response.html(rendered_template)

@bp.route('/<zone>/submit-spots')
async def submit_spots(request, zone):
    tenant = TenantManagement(123)
    zoneInstance = ZoneManagement(zone)
    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_list='true', 
        zoneName=zone,
        tenantName=tenant.name,
        zonePolygon=zoneInstance.getPolygon(),
        spotList=zoneInstance.getSpotList()
    )
    # checking spot list and adding to DB
    return response.html(rendered_template)

# Handling Parking Spots
@bp.route('/<zone>/spots')
async def spots(request, zone):
    tenant = TenantManagement(123)
    zoneInstance = ZoneManagement(zone)
    spotsJson = jsonList(zoneInstance.getSpotList())
    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_list='true', 
        zoneName=zone,
        tenantName=tenant.name,
        zonePolygon=zoneInstance.getPolygon(),
        spotList = spotsJson
    )
    return response.html(rendered_template)

# convert elements of list into Json
def jsonList(arg):
    liste = []
    for item in arg:
        liste.append(js.dumps(item.toJson()))
    return liste
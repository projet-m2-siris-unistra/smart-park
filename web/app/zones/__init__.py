import json as js

from sanic import Blueprint, response
from sanic.response import json

from app.templating import render

from app.parkings import Tooling
from app.parkings import ZoneManagement
from app.parkings import TenantManagement

from app.bus import Request

bp = Blueprint("zones", url_prefix='/parking/zone')

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

@bp.route('/<zone_id>')
@bp.route('/<zone_id>/overview')
async def view(request, zone_id):
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)

    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_view='true', 
        zoneName=zoneInstance.name,
        zone_id=zone_id,
        tenantName=tenantInstance.name,
        zoneTaken=zoneInstance.getNbTakenSpots(),
        zoneTotal=zoneInstance.getNbTotalSpots()
    )
    return response.html(rendered_template)

@bp.route('/<zone_id>/statistics')
async def stats(request, zone_id):
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)
   
    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)

    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_stats='true', 
        zoneName=zoneInstance.name,
        zone_id=zone_id,
        tenantName=tenantInstance.name,
        statistics=zoneInstance.getAllStats()
    )
    return response.html(rendered_template)

@bp.route('/<zone_id>/maintenance')
async def maintenance(request, zone_id):
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)

    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_maintenance='true',
        zoneName=zoneInstance.name,
        zone_id=zone_id,
        tenantName=tenantInstance.name,    
        spotList=zoneInstance.getSpotList()
    )
    return response.html(rendered_template)

@bp.route('/<zone_id>/configuration')
async def config(request, zone_id):
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)

    await zoneInstance.getSpotList()
    spotsJson = Tooling.jsonList(zoneInstance.spots)

    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_config='true',
        zoneName=zoneInstance.name,
        tenantName=tenantInstance.name,
        zone_id=zone_id,
        zonePolygon=zoneInstance.polygon,
        zoneInstance=zoneInstance,
        zoneColor=zoneInstance.color,
        tenantCoor=tenantInstance.coordinates,
        spotList=spotsJson
    )
    return response.html(rendered_template)

@bp.route('/<zone_id>/submit-spots')
async def submit_spots(request, zone_id):
    tenant = TenantManagement(1)
    await tenant.init(1)
    
    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)

    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_list='true', 
        zoneName=zoneInstance.name,
        zone_id=zone_id,
        tenantName=tenant.name,
        zonePolygon=zoneInstance.polygon,
        spotList=zoneInstance.getSpotList()
    )
    # checking spot list and adding to DB
    return response.html(rendered_template)

# Handling Parking Spots
@bp.route('/<zone_id>/spots')
async def spots(request, zone_id):
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)

    await zoneInstance.getSpotList()
    spotsJson = Tooling.jsonList(zoneInstance.spots)
    print("spotJson: ", spotsJson)
    print("spots: ", zoneInstance.spots)

    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_list='true', 
        zoneName=zoneInstance.name,
        tenantName=tenantInstance.name,
        zone_id=zone_id,
        tenantCoor=tenantInstance.coordinates,
        zonePolygon=zoneInstance.polygon,
        zoneColor=zoneInstance.color,
        spotList=spotsJson
    )
    return response.html(rendered_template)

# Interface for deleting a zone
@bp.route('/<zone_id>/remove')
async def remove(request, zone_id):
    tenant = TenantManagement(1)
    await tenant.init(1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)
    
    rendered_template = await render(
        'zone_removing.html', 
        request,
        zoneName=zoneInstance.name,
        tenantName=tenant.name
    )
    return response.html(rendered_template)

@bp.route('/<zone_id>/remove-check')
async def remove_check(request, zone_id):
    # removing the zone
    # removing the spots if needed

    # temporary redirection to dashboard of tenant
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)
    rendered_template = await render(
        "dashboard_template.html",
        request, 
        zoneList=tenantInstance.getZones(),
        totalSpots=tenantInstance.getTotalSpots(),
        takenSpots=tenantInstance.getTakenSpots(),
        removed_zone=True
    )
    return response.html(rendered_template)

from sanic import Blueprint, response
from sanic.response import json

from app.templating import render

from app.parkings import ZoneManagement

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
    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_view='true', 
        zoneName=zone
    )
    return response.html(rendered_template)

@bp.route('/<zone>/statistics')
async def stats(request, zone):
    zoneInstance = ZoneManagement(zone)
    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_stats='true', 
        zoneName=zone,
        statistics=zone1.getAllStats()
    )
    return response.html(rendered_template)

@bp.route('/<zone>/configuration')
async def config(request, zone):
    zoneInstance = ZoneManagement(zone)
    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_config='true',
        zoneName=zone,
        zonePolygon=zoneInstance.getPolygon(),
        spotList=zoneInstance.getSpotList()
    )
    return response.html(rendered_template)

@bp.route('/<zone>/submit-spots')
async def submit_spots(request, zone):
    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_list='true', 
        zoneName=zone
    )
    # checking spot list and adding to DB
    return response.html(rendered_template)


# Handling Parking Spots
@bp.route('/<zone>/spots')
async def spots(request, zone):
    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_list='true', 
        zonePolygon=zoneInstance.getPolygon(),
        zone=zone
    )
    return response.html(rendered_template)

@bp.route('/<zone>/spots/<spot>')
async def spot_edit(request, zone, spot):
    return response.text("edition de la place")

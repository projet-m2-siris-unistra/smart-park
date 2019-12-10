import json as js

from sanic import Blueprint, response
from sanic.response import json
from sanic.exceptions import ServerError

from app.templating import render

from app.parkings import Tooling
from app.parkings import ZoneManagement
from app.parkings import TenantManagement

from app.bus import Request

from app.forms.zones import CreationForm

bp = Blueprint("zones", url_prefix='/parking/zone')



# Handling Parkings zones
@bp.route('/create_zone', methods=['POST', 'GET'])
async def create_zone(request):
    form = CreationForm(request)

    name = form.name.data
    print("Name=", name)
    type = form.type.data
    print("Type=", type)
    color = form.color.data
    print("Color=", color)
    polygon = form.polygon.data
    print("Polygon=", polygon)

    if form.validate_on_submit():
        print("Form validated")
        await Request.createZone(
            tenant_id=1, 
            name=name, 
            type=type, 
            color=color[1:].upper(), # cut the '#' and upper letters 
            polygon=polygon
        )
        return response.redirect('/dashboard')

    rendered_template = await render(
        "zone_creation.html", 
        request,
        form=form
    )
    return response.html(rendered_template)


@bp.route('/<zone_id>')
@bp.route('/<zone_id>/overview')
async def view(request, zone_id):
    tenantInstance = TenantManagement(tenant_id=1)
    await tenantInstance.init(tenant_id=1)

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
    tenantInstance = TenantManagement(tenant_id=1)
    await tenantInstance.init(tenant_id=1)
   
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
    tenantInstance = TenantManagement(tenant_id=1)
    await tenantInstance.init(tenant_id=1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)

    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_maintenance='true',
        zoneName=zoneInstance.name,
        zone_id=zone_id,
        tenantName=tenantInstance.name,    
        spotList=zoneInstance.setSpots()
    )
    return response.html(rendered_template)


@bp.route('/<zone_id>/configuration', methods=['POST', 'GET'])
async def config(request, zone_id):
    tenantInstance = TenantManagement(tenant_id=1)
    await tenantInstance.init(tenant_id=1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)

    await zoneInstance.setSpots()
    spotsJson = Tooling.jsonList(zoneInstance.spots)

    if request.method == 'POST':
        print("Form validated")
        print("args: ", request.args)
        print("raw_args: ", request.raw_args)
        #return response.redirect()

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
    tenant = TenantManagement(tenant_id=1)
    await tenant.init(tenant_id=1)
    
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
        spotList=zoneInstance.setSpots()
    )
    # checking spot list and adding to DB
    return response.html(rendered_template)


# Handling Parking Spots
@bp.route('/<zone_id>/spots')
async def spots(request, zone_id):
    tenantInstance = TenantManagement(tenant_id=1)
    await tenantInstance.init(tenant_id=1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)

    await zoneInstance.setSpots()
    spotsJson = Tooling.jsonList(zoneInstance.spots)

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
    tenantInstance = TenantManagement(tenant_id=1)
    await tenantInstance.init(tenant_id=1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)
    
    rendered_template = await render(
        'zone_removing.html', 
        request,
        zone_id=zone_id,
        zoneName=zoneInstance.name,
        tenantName=tenantInstance.name
    )
    return response.html(rendered_template)


@bp.route('/<zone_id>/remove-check')
async def remove_check(request, zone_id):
    # removing the zone
    # removing the spots if needed

    # temporary redirection to dashboard of tenant
    tenantInstance = TenantManagement(tenant_id=1)
    await tenantInstance.init(tenant_id=1)

    await tenantInstance.setZones()
    if tenantInstance.zones is None:
        raise ServerError("No zones in DB", status_code=500)

    rendered_template = await render(
        "dashboard_template.html",
        request, 
        zoneList=tenantInstance.zones,
        totalSpots=tenantInstance.getTotalSpots(),
        takenSpots=tenantInstance.getTakenSpots(),
        removed_zone=True
    )
    return response.html(rendered_template)

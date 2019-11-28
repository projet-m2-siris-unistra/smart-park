from sanic import Blueprint, response
from sanic.response import json

from app.parkings import TenantManagement
from app.parkings import ZoneManagement
from app.parkings import SpotManagement

from app.templating import render


bp = Blueprint("spots", url_prefix='/parking/zone/<zone_id>/spot')

@bp.route('/<spot_id>')
@bp.route('/<spot_id>/overview')
async def overview(request, spot_id, zone_id):
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)

    spotInstance = SpotManagement(spot_id)
    await spotInstance.init(spot_id)

    rendered_template = await render(
        'spot_template.html', 
        request,
        active_tab_overview='true', 
        zone_id=zone_id,
        spot_id=spot_id,
        spotName=spotInstance.name,
        zoneName=zoneInstance.name,
        tenantName=tenantInstance.name
    )
    return response.html(rendered_template)


@bp.route('/<spot_id>/statistics')
async def stats(request, spot_id, zone_id):
    tenantInstance = TenantManagement(123)
    await tenantInstance.init(1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)

    spotInstance = SpotManagement(spot_id)
    await spotInstance.init(spot_id)

    rendered_template = await render(
        'spot_template.html', 
        request,
        active_tab_stats='true', 
        spotName=spotInstance.name,
        zone_id=zone_id,
        spot_id=spot_id,
        zoneName=zoneInstance.name,
        statistics=spotInstance.getAllStats(),
        tenantName=tenantInstance.name
    )
    return response.html(rendered_template)


@bp.route('/<spot_id>/maintenance')
async def maintenance(request, spot_id, zone_id):
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)

    spotInstance = SpotManagement(spot_id)
    await spotInstance.init(spot_id)

    rendered_template = await render(
        'spot_template.html', 
        request,
        active_tab_maintenance='true', 
        zone_id=zone_id,
        spot_id=spot_id,
        spotName=spotInstance.name,
        zoneName=zoneInstance.name,
        tenantName=tenantInstance.name
    )
    return response.html(rendered_template)


@bp.route('/<spot_id>/configuration')
async def config(request, spot_id, zone_id):
    tenantInstance = TenantManagement(1)
    await tenantInstance.init(1)
    
    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)

    spotInstance = SpotManagement(spot_id)
    await spotInstance.init(spot_id)

    rendered_template = await render(
        'spot_template.html', 
        request,
        active_tab_config='true', 
        zone_id=zone_id,
        spot_id=spot_id,
        spotName=spotInstance.name,
        zoneName=zoneInstance.name,
        tenantName=tenantInstance.name
    )
    return response.html(rendered_template)
from sanic import Blueprint, response
from sanic.response import json

from app.parkings import TenantManagement
from app.parkings import ZoneManagement
from app.parkings import SpotManagement

from app.templating import render


bp = Blueprint("spots", url_prefix='/spot')

@bp.route('/<spot>')
@bp.route('/<spot>/overview')
async def overview(request, spot):
    tenant = TenantManagement(123)
    zone = ZoneManagement("CENTRE")
    rendered_template = await render(
        'spot_template.html', 
        request,
        active_tab_overview='true', 
        spotName=spot,
        zoneName=zone.name,
        tenantName=tenant.name
    )
    return response.html(rendered_template)

@bp.route('/<spot>/statistics')
async def stats(request, spot):
    tenantInstance = TenantManagement(123)
    zoneInstance = ZoneManagement("CENTRE")
    spotInstance = SpotManagement()
    rendered_template = await render(
        'spot_template.html', 
        request,
        active_tab_stats='true', 
        spotName=spot,
        zoneName=zoneInstance.name,
        statistics=spotInstance.getAllStats(),
        tenantName=tenantInstance.name
    )
    return response.html(rendered_template)

@bp.route('/<spot>/maintenance')
async def maintenance(request, spot):
    tenant = TenantManagement(123)
    zone = ZoneManagement("CENTRE")
    rendered_template = await render(
        'spot_template.html', 
        request,
        active_tab_maintenance='true', 
        spotName=spot,
        zoneName=zone.name,
        tenantName=tenant.name
    )
    return response.html(rendered_template)

@bp.route('/<spot>/configuration')
async def config(request, spot):
    tenant = TenantManagement(123)
    zone = ZoneManagement("CENTRE")
    rendered_template = await render(
        'spot_template.html', 
        request,
        active_tab_config='true', 
        spotName=spot,
        zoneName=zone.name,
        tenantName=tenant.name
    )
    return response.html(rendered_template)
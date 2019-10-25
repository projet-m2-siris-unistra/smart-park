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


@bp.route('/<zone>')
@bp.route('/<zone>/overview')
async def view(request, zone):
    rendered_template = await render('parking_template.html', request,
        active_tab_view='true', zone=zone)
    return response.html(rendered_template)

@bp.route('/<zone>/spots')
async def spots(request, zone):
    rendered_template = await render('parking_template.html', request,
        active_tab_list='true', zone=zone)
    return response.html(rendered_template)

@bp.route('/<zone>/statistics')
async def stats(request, zone):
    zone1 = ZoneManagement()
    rendered_template = await render('parking_template.html', request,
        active_tab_stats='true', 
        zone=zone,
        statistics=zone1.getAllStats()
        )
    return response.html(rendered_template)

@bp.route('/<zone>/configuration')
async def config(request, zone):
    rendered_template = await render('parking_template.html', request,
        active_tab_config='true', zone=zone)
    return response.html(rendered_template)

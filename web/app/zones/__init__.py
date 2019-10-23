from sanic import Blueprint, response
from sanic.response import json

from app.templating import render

from app.parkings import ZoneManagement

bp = Blueprint(__name__, url_prefix='/parking')

# Handling Parkings zones
@bp.route('/<zone>')
@bp.route('/<zone>/overview')
async def parking_view(request, zone):
    rendered_template = await render('parking_template.html',
        active_tab_view='true', zone=zone)
    return response.html(rendered_template)

@bp.route('/<zone>/spots')
async def parking_spots(request, zone):
    rendered_template = await render('parking_template.html',
        active_tab_list='true', zone=zone)
    return response.html(rendered_template)

@bp.route('/<zone>/statistics')
async def parking_stats(request, zone):
    zone1 = ZoneManagement()
    rendered_template = await render('parking_template.html',
        active_tab_stats='true', 
        zone=zone,
        statistics=zone1.getAllStats()
        )
    return response.html(rendered_template)

@bp.route('/<zone>/configuration')
async def parking_config(request, zone):
    rendered_template = await render('parking_template.html',
        active_tab_config='true', zone=zone)
    return response.html(rendered_template)

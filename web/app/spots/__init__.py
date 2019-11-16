from sanic import Blueprint, response
from sanic.response import json

from app.templating import render

from app.parkings import ZoneManagement

bp = Blueprint("spot", url_prefix='/parking/<zone>')

@bp.route('/<spot>')
@bp.route('/<spot>/overview')
async def overview(request, spot):
    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_overview='true', 
        spotName=spot
    )
    return response.html(rendered_template)
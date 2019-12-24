from sanic import Blueprint, response
from sanic.response import json

from app.parkings import TenantManagement
from app.parkings import ZoneManagement
from app.parkings import SpotManagement
from app.forms.spots import ConfigurationForm
from app.bus import Request
from app import ServerError

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
        tenantInstance=tenantInstance,
        zoneInstance=zoneInstance,
        spotInstance=spotInstance
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
        zone_id=zone_id,
        spot_id=spot_id,
        tenantInstance=tenantInstance,
        zoneInstance=zoneInstance,
        spotInstance=spotInstance,
        statistics=spotInstance.getAllStats(),
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
        tenantInstance=tenantInstance,
        zoneInstance=zoneInstance,
        spotInstance=spotInstance
    )
    return response.html(rendered_template)


@bp.route('/<spot_id>/configuration', methods=['GET', 'POST'])
async def config(request, spot_id, zone_id):
    tenantInstance = TenantManagement(1)
    res = await tenantInstance.init(1)
    if res == Request.REQ_ERROR:
        raise ServerError("impossible de récupérer le tenant", status_code=500)

    zoneInstance = ZoneManagement(zone_id)
    res = await zoneInstance.init(zone_id)
    if res == Request.REQ_ERROR:
        raise ServerError("impossible de récupérer la zone", status_code=500)

    spotInstance = SpotManagement(spot_id)
    res = await spotInstance.init(spot_id)
    if res == Request.REQ_ERROR:
        raise ServerError("impossible de récupérer la place", status_code=500)
    
    res = await spotInstance.setDevice()
    if res == Request.REQ_ERROR:
        raise ServerError("impossible de récupérer le capteur", status_code=500)

    # Creating and filing formular
    form = ConfigurationForm(request, spotInstance)

    # performing adaptation of coordinates string (remove '[' and ']')
    if '[' and ']' in form.coordinates.data:
        form.coordinates.data = form.coordinates.data[1:len(form.coordinates.data)-2]

    # Making the device select choices
    deviceList = []
    deviceList.append((spotInstance.device.id, spotInstance.device.eui))

    res = await tenantInstance.getNotAssignedDevices()
    if res == Request.REQ_ERROR:
        raise ServerError("Impossible de charger les capteurs non assignés.")
    deviceList = deviceList + res
    form.device.choices = deviceList

    # Form validation check & request
    if form.validate_on_submit():
        
        # The actual update request to the queue
        if form.submit.data:
            print("Submit validated")
            lnglat = form.coordinates.data.split(',')
            res = await Request.updateSpot(
                spot_id=spotInstance.id,
                device_id=form.device.data,
                type=form.type.data,
                coordinates=[float(lnglat[0]) ,float(lnglat[1])]
            )
            if res == Request.REQ_ERROR:
                raise ServerError("Impossible de mettre à jour la place.")
            return response.redirect("/parking/zone/"+zone_id+"/spot/"+spot_id)

        # Handling place deletion
        elif form.delete.data:
            print("spot deletion")
            res = await Request.deleteSpot(spotInstance.id)
            if res == Request.REQ_ERROR:
                raise ServerError("impossible de supprimer la place.")
            return response.redirect("/parking/zone/"+zone_id)

    rendered_template = await render(
        'spot_template.html', 
        request,
        active_tab_config='true', 
        zone_id=zone_id,
        spot_id=spot_id,
        tenantInstance=tenantInstance,
        zoneInstance=zoneInstance,
        spotInstance=spotInstance,
        form=form
    )
    return response.html(rendered_template)
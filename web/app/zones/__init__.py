import json as js

from sanic import Blueprint, response
from sanic.response import json
from sanic.exceptions import ServerError

from app.templating import render

from app.parkings import Tooling
from app.parkings import ZoneManagement
from app.parkings import TenantManagement

from app.bus import Request

from app.forms.zones import CreationForm, ConfigurationForm, SpotsAddingForm

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
            color=Tooling.formatColor(color),
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
        zone_id=zone_id,
        zoneInstance=zoneInstance,
        tenantInstance=tenantInstance
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
        zone_id=zone_id,
        zoneInstance=zoneInstance,
        tenantInstance=tenantInstance,
        statistics=zoneInstance.getAllStats()
    )
    return response.html(rendered_template)


@bp.route('/<zone_id>/maintenance')
async def maintenance(request, zone_id):
    tenantInstance = TenantManagement(tenant_id=1)
    await tenantInstance.init(tenant_id=1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)
    await zoneInstance.setSpots()

    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_maintenance='true',
        zone_id=zone_id,
        zoneInstance=zoneInstance,
        tenantInstance=tenantInstance
    )
    return response.html(rendered_template)


@bp.route('/<zone_id>/configuration', methods=['POST', 'GET'])
async def config(request, zone_id):
    changes = False
    
    tenantInstance = TenantManagement(tenant_id=1)
    await tenantInstance.init(tenant_id=1)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)
    await zoneInstance.setSpots()
    
    formGeneral = ConfigurationForm(request, zoneInstance)

    # We need the not assigned devices from this tenant
    deviceList = await tenantInstance.getNotAssignedDevices()
    if deviceList == Request.REQ_ERROR:
        raise ServerError("Impossible de charger les capteurs non assignés.")

    formSpots = SpotsAddingForm(request)

    print("delete data:", formGeneral.delete.data)
    print("submit data: ", formGeneral.submit.data)
    print("name data: ", formGeneral.name.data)

    # Handling general configuration
    if formGeneral.validate_on_submit():

        print("delete data:", formGeneral.delete.data)
        print("submit data: ", formGeneral.submit.data)

        if formGeneral.delete.data:
            print("Zone deletion")

        elif formGeneral.submit.data:
            print("Form validated")
            changes = True
            await Request.updateZone(
                zone_id=zone_id,
                tenant_id=1, 
                name=formGeneral.name.data, 
                type=formGeneral.type.data, 
                color=Tooling.formatColor(formGeneral.color.data),
                polygon=""
            )

        else:
            print("WARNING: Wrong button clicked")

    # handling spot adding form
    if formSpots.validate_on_submit():
        print("spots adding form validated")

    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_config='true',
        zone_id=zone_id,
        zoneInstance=zoneInstance,
        tenantInstance=tenantInstance,
        formGeneral=formGeneral,
        formSpots=formSpots,
        spotList=Tooling.jsonList(zoneInstance.spots),
        changesApplied=changes
    )
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
        zone_id=zone_id,
        zoneInstance=zoneInstance,
        tenantInstance=tenantInstance,
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
        zoneInstance=zoneInstance,
        tenantInstance=tenantInstance
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
        tenantInstance=tenantInstance,
        removed_zone=True
    )
    return response.html(rendered_template)

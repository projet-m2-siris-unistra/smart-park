import json as js

from sanic import Blueprint, response
from sanic.response import json
from sanic.exceptions import ServerError

from app.templating import render

from app.parkings import Tooling
from app.parkings import ZoneManagement
from app.parkings import TenantManagement
from app.pagination import Pagination

from app.bus import Request

from app.forms.zones import CreationForm, ConfigurationForm, SpotsAddingForm

bp = Blueprint("zones", url_prefix='/parking/zone')


# Checking if the user is allowed to display the ressource
@bp.middleware('request')
async def check_session(request):
    print("Session checking...")
    session_cookie = request.cookies.get('session')


# Handling Parkings zones
@bp.route('/create_zone', methods=['POST', 'GET'])
async def create_zone(request):
    form = CreationForm(request)

    if form.validate_on_submit():
        print("Form validated")
        await Request.createZone(
            tenant_id=request.ctx.tenant_id,
            name=form.name.data, 
            type=form.type.data, 
            color=Tooling.formatColor(form.color.data),
            polygon=form.polygon.data
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
    tenantInstance = TenantManagement(request.ctx.tenant_id)
    await tenantInstance.init(request.ctx.tenant_id)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)
    await zoneInstance.setSpots()

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
    tenantInstance = TenantManagement(request.ctx.tenant_id)
    await tenantInstance.init(request.ctx.tenant_id)
   
    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)
    statistics = zoneInstance.getAllStats()

    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_stats='true',
        zone_id=zone_id,
        zoneInstance=zoneInstance,
        tenantInstance=tenantInstance,
        statistics=statistics
    )
    return response.html(rendered_template)


@bp.route('/<zone_id>/maintenance')
async def maintenance(request, zone_id):
    tenantInstance = TenantManagement(request.ctx.tenant_id)
    await tenantInstance.init(request.ctx.tenant_id)

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
    
    tenantInstance = TenantManagement(request.ctx.tenant_id)
    await tenantInstance.init(request.ctx.tenant_id)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)
    await zoneInstance.setSpots()
    
    formGeneral = ConfigurationForm(request, zoneInstance, prefix="genForm")
    formSpots = SpotsAddingForm(request, prefix="spotsForm")

    # We need the not assigned devices from this tenant
    deviceList = await tenantInstance.getNotAssignedDevices()
    if deviceList == Request.REQ_ERROR:
        raise ServerError("Impossible de charger les capteurs non assignés.")
    
    # Creating form and setting device selection
    formSpots.deviceSelect.choices = deviceList
    
    # Informing user if no device is available
    if deviceList == []:
        formSpots.deviceSelect.description = "ATTENTION : aucun capteur disponible !"

    # Handling general configuration
    if formGeneral.validate_on_submit():

        # Deletion clicked 
        if formGeneral.delete.data:
            print("Zone deletion")
            res = await Request.deleteZone(zoneInstance.id)
            if res == Request.REQ_ERROR:
                raise ServerError("erreur lors de la suppression de la zone")
            else:
                print("RES: ", res)
                return response.redirect("/dashboard")

        # General form validated
        elif formGeneral.submitGen.data:
            print("General form validated")
            res = await Request.updateZone(
                zone_id=zone_id,
                tenant_id=request.ctx.tenant_id, 
                name=formGeneral.name.data, 
                type=formGeneral.type.data, 
                color=Tooling.formatColor(formGeneral.color.data),
                polygon=""
            )
            if res == Request.REQ_ERROR:
                raise ServerError("impossible de mettre à jour la zone", 500)
            changes = True
            

    # Checking if user wants to add spots
    if formSpots.validate_on_submit():
        print("Spot adding form validated")
        lnglat = formSpots.coordinatesInput.data.split(',')

        res = await Request.createSpot(
            zone_id=zone_id,
            device_id=formSpots.deviceSelect.data,
            type=formSpots.typeSelect.data,
            coordinates=[float(lnglat[0]) ,float(lnglat[1])]
        )
        if res == Request.REQ_ERROR:
            raise ServerError("impossible d'ajouter une place", 500)
        return response.redirect("/parking/zone/"+zone_id+"/spots")

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
        
    tenantInstance = TenantManagement(request.ctx.tenant_id)
    await tenantInstance.init(request.ctx.tenant_id)

    # Calculating and initialization of pagination
    pagination = Pagination(request)

    zoneInstance = ZoneManagement(zone_id)
    await zoneInstance.init(zone_id)
    await zoneInstance.setSpots(
        page=pagination.page_number, 
        pagesize=pagination.page_size
    )

    # Handling request and tab
    tab_map = True

    if request.raw_args.get('type') == "table":
        tab_map = False
        for spot in zoneInstance.spots:
            await spot.setDevice()

    pagination.setElementsNumber(zoneInstance.spotsCount)

    spotsJson = Tooling.jsonList(zoneInstance.spots)

    rendered_template = await render(
        'parking_template.html', 
        request,
        active_tab_list='true', 
        zone_id=zone_id,
        zoneInstance=zoneInstance,
        tenantInstance=tenantInstance,
        spotList=spotsJson,
        pagination = pagination,
        tab_map=tab_map
    )
    return response.html(rendered_template)


# Interface for deleting a zone
@bp.route('/<zone_id>/remove')
async def remove(request, zone_id):
    tenantInstance = TenantManagement(request.ctx.tenant_id)
    await tenantInstance.init(request.ctx.tenant_id)

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
    tenantInstance = TenantManagement(request.ctx.tenant_id)
    await tenantInstance.init(request.ctx.tenant_id)

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

import json as js

from sanic import Blueprint, response
from sanic.response import json
from sanic.exceptions import ServerError

from app.templating import render

from app.parkings import Tooling
from app.parkings import TenantManagement
from app.parkings import DeviceManagement
from app.pagination import Pagination

from app.bus import Request

from app.forms.devices import CreationForm

bp = Blueprint("devices", url_prefix='/devices')



# List of devices
@bp.route("/list", methods=['GET', 'POST'])
async def view(request):
    tenantInstance = TenantManagement(tenant_id=1)
    await tenantInstance.init(tenant_id=1)

    pagination = Pagination(request)
    await tenantInstance.setDevices(
        page=pagination.page_number, 
        pagesize=pagination.page_size
    )
    pagination.setElementsNumber(tenantInstance.devicesCount)

    # User wants to delete the device
    if request.method == 'POST':
        print("Deletion of device_id : ", request.form.get('device-id'))
        res = await Request.deleteDevice(request.form.get('device-id'))
        if res == Request.REQ_ERROR:
            raise ServerError("impossible de supprimer le device")
        return response.redirect("/dashboard")

    rendered_template = await render(
        "devices_template.html", 
        request,
        devices=tenantInstance.devices,
        pagination=pagination
    )
    return response.html(rendered_template)


# Creating a device
@bp.route('/create_device', methods=['POST', 'GET'])
async def create(request):
    form = CreationForm(request)

    tenantInstance = TenantManagement(tenant_id=1)
    await tenantInstance.init(tenant_id=1)

    if form.validate_on_submit():
        print("Form validated")
        print("eui=", form.eui.data)
        print("name=", form.name.data)
        res = await Request.createDevice(
            tenant_id=1,
            eui=form.eui.data,
            name=form.name.data
        )

        # Checking if request worked
        if res == Request.REQ_ERROR:
            raise ServerError("Cannot create device", status_code=500)
        else:
            return response.redirect("/devices/list")


    rendered_template = await render(
        "devices_creation.html", 
        request,
        form=form
    )
    return response.html(rendered_template)
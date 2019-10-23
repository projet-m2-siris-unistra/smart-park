from sanic import Sanic
from sanic import response
from sanic.response import json
from jinja2 import Environment, PackageLoader, select_autoescape

import app.parkings
from app.templating import render


app = Sanic(__name__)


app.static("/static", "./static")



# Handling navigation
@app.route('/')
@app.route('/home')
async def welcome(request):
    rendered_template = await render('base_template.html',
        knights='BIENVENUE SUR SMART PARK !')
    return response.html(rendered_template)

@app.route('/dashboard')
async def dashboard(request):
    rendered_template = await render('dashboard_template.html')
    return response.html(rendered_template)

@app.route('/zones')
async def zones(request):
    rendered_template = await render('zones_template.html')
    return response.html(rendered_template)



# Handling accounts
@app.route('/login')
async def login(request):
    rendered_template = await render('login_template.html')
    return response.html(rendered_template)

@app.route('/login-check', methods=["POST"])
async def login_check(request):
    return json({"form": request.form})

@app.route('/signup-check', methods=["POST"])
async def login_check(request):
    return json({"form": request.form})

@app.route('/signup')
async def login(request):
    rendered_template = await render('signup_template.html')
    return response.html(rendered_template)



# Handling Parkings zones
@app.route('/parking/<zone>')
@app.route('/parking/<zone>/overview')
async def parking_view(request, zone):
    rendered_template = await render('parking_template.html',
        active_tab_view='true', zone=zone)
    return response.html(rendered_template)

@app.route('/parking/<zone>/spots')
async def parking_spots(request, zone):
    rendered_template = await render('parking_template.html',
        active_tab_list='true', zone=zone)
    return response.html(rendered_template)

@app.route('/parking/<zone>/statistics')
async def parking_stats(request, zone):
    zone1 = parkings.ZoneManagement()
    rendered_template = await render('parking_template.html',
        active_tab_stats='true', 
        zone=zone,
        statistics=zone1.getAllStats()
        )
    return response.html(rendered_template)

@app.route('/parking/<zone>/configuration')
async def parking_config(request, zone):
    rendered_template = await render('parking_template.html',
        active_tab_config='true', zone=zone)
    return response.html(rendered_template)

def run():
    app.run(host="0.0.0.0", port=8080, debug=True)

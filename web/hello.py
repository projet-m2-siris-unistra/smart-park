from sanic import Sanic
from sanic import response
from sanic.response import json
from jinja2 import Environment, PackageLoader, select_autoescape
import parkings


app = Sanic(__name__)

# Load the template environment with async support
template_env = Environment(
    loader=PackageLoader('hello', 'templates'),
    autoescape=select_autoescape(['html', 'xml']),
    enable_async=True
)

# Load the template from file
base_template = template_env.get_template("base_template.html")
dashboard_template = template_env.get_template("dashboard_template.html")
zones_template = template_env.get_template("zones_template.html")
login_template = template_env.get_template("login_template.html")
signup_template = template_env.get_template("signup_template.html")
# view of a parking zone, with the listing of all places
parking_template = template_env.get_template("parking_template.html")


app.static("/static", "./static")



# Handling navigation
@app.route('/')
@app.route('/home')
async def welcome(request):
    rendered_template = await base_template.render_async(
        knights='BIENVENUE SUR SMART PARK !')
    return response.html(rendered_template)

@app.route('/dashboard')
async def dashboard(request):
    rendered_template = await dashboard_template.render_async()
    return response.html(rendered_template)

@app.route('/zones')
async def zones(request):
    rendered_template = await zones_template.render_async()
    return response.html(rendered_template)



# Handling accounts
@app.route('/login')
async def login(request):
    rendered_template = await login_template.render_async()
    return response.html(rendered_template)

@app.route('/login-check', methods=["POST"])
async def login_check(request):
    return json({"form": request.form})

@app.route('/signup-check', methods=["POST"])
async def login_check(request):
    return json({"form": request.form})

@app.route('/signup')
async def login(request):
    rendered_template = await signup_template.render_async()
    return response.html(rendered_template)



# Handling Parkings zones
@app.route('/parking/<zone>')
@app.route('/parking/<zone>/overview')
async def parking_view(request, zone):
    rendered_template = await parking_template.render_async(
        active_tab_view='true', zone=zone)
    return response.html(rendered_template)

@app.route('/parking/<zone>/spots')
async def parking_spots(request, zone):
    rendered_template = await parking_template.render_async(
        active_tab_list='true', zone=zone)
    return response.html(rendered_template)

@app.route('/parking/<zone>/statistics')
async def parking_stats(request, zone):
    zone1 = parkings.ZoneManagement()
    rendered_template = await parking_template.render_async(
        active_tab_stats='true', 
        zone=zone,
        statistics=zone1.getAllStats()
        )
    return response.html(rendered_template)

@app.route('/parking/<zone>/configuration')
async def parking_config(request, zone):
    rendered_template = await parking_template.render_async(
        active_tab_config='true', zone=zone)
    return response.html(rendered_template)



# starting website
app.run(host="0.0.0.0", port=8080, debug=True)

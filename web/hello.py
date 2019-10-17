from sanic import Sanic
from sanic import response
from jinja2 import Environment, PackageLoader, select_autoescape

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

app.static("/static", "./static")

@app.route('/')
@app.route('/acceuil')
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

@app.route('/connexion')
async def login(request):
    rendered_template = await login_template.render_async()
    return response.html(rendered_template)

@app.route('/login-check', methods=['POST'])
async def login_check(request):
    return response.text(request.args)

app.run(host="0.0.0.0", port=8080, debug=True)
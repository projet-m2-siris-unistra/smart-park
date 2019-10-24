import click
from sanic import Sanic
from sanic import response
from sanic.response import json
from jinja2 import Environment, PackageLoader, select_autoescape

import app.config
import app.accounts
import app.zones
from app.templating import render


app = Sanic(__name__)
app.register_listener(config.load, 'before_server_start')

app.static("/static", "./static")

app.blueprint(accounts.bp)
app.blueprint(zones.bp)

# Handling navigation
@app.route('/')
@app.route('/home')
async def home(request):
    rendered_template = await render('base_template.html', request,
        knights='BIENVENUE SUR SMART PARK !')
    return response.html(rendered_template)

@app.route('/dashboard')
async def dashboard(request):
    rendered_template = await render('dashboard_template.html', request)
    return response.html(rendered_template)

@app.route('/zones')
async def zones(request):
    rendered_template = await render('zones_template.html', request)
    return response.html(rendered_template)

@click.command()
@config.run_params
def run():
    app.run(host=config.env('HOST'),
            port=config.env('PORT'),
            debug=config.env('DEBUG'))

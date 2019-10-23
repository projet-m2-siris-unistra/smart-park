from sanic import Blueprint, response
from sanic.response import json

from app.templating import render

bp = Blueprint(__name__)

@bp.route('/login')
async def login(request):
    rendered_template = await render('login_template.html')
    return response.html(rendered_template)

@bp.route('/login-check', methods=["POST"])
async def login_check(request):
    return json({"form": request.form})

@bp.route('/signup-check', methods=["POST"])
async def login_check(request):
    return json({"form": request.form})

@bp.route('/signup')
async def login(request):
    rendered_template = await render('signup_template.html')
    return response.html(rendered_template)


from sanic import Blueprint, response, cookies
from sanic.response import json
from sanic.exceptions import ServerError

import json as js

from app.templating import render
from app.bus import Auth
from app.bus.Request import REQ_ERROR

bp = Blueprint("accounts")



async def whatever(request):
    if request.cookies.get('session'):
        cookie = request.cookies.get('session')

        # Validate token
        res = await Auth.validate(cookie)
        if res == REQ_ERROR:
            return False

        # Get Session infos
        data = js.loads(res)
        if not data['valid']:
            return False

        request.ctx.tenant_id = data['tenant_id']
        request.ctx.user = data['user']
        return True


# List of urls with no session needed
offlinePaths = [
    "/",
    "/auth",
    "/home",
    "/about",
    "/faq",
    "/logout"
]


# Handling session
@bp.middleware('request')
async def session_authorization(request):
    
    # checking if a session cookie already exists
    if await whatever(request):
        return

    # Ignore not protected pages
    if request.path in offlinePaths or "/static/" in request.path:
        return 

    else:
        # Redirecting to SSO
        print("Redirection to SSO...")
        url = request.url_for("accounts.session_exchange")
        res = await Auth.authorization(request.host, url)
        data = js.loads(res)
        resp = response.redirect(data['url'])
        del resp.cookies["session"]
        return resp


@bp.route('/auth')
async def session_exchange(request):

    if request.raw_args.get('code'):
        code = request.raw_args.get('code')

        # Check the code to get session token
        url = request.url_for("accounts.session_exchange")
        res = await Auth.exchange(request.host, code, url)
        if res == REQ_ERROR:
            return response.text("IMPOSSIBLE DE SE CONNECTER")

        data = js.loads(res)
        token = data['token']

        resp = response.redirect("/")
        resp.cookies["session"] = token

        return resp

    else:
        return response.text("IMPOSSIBLE DE SE CONNECTER")


@bp.route('/login')
async def login(request):
    if request.cookies.get('session'):
        return response.redirect("/")

    # Redirecting user to his SSO
    print("Login started...")


@bp.route('/logout')
async def logout(request):
    # delete the cookie
    if request.cookies.get('session'):
        resp = response.redirect("/")
        del resp.cookies['session']
        return resp

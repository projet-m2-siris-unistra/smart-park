from jinja2 import Environment, PackageLoader, select_autoescape

env = Environment(
    loader=PackageLoader("app", "templates"),
    autoescape=select_autoescape(['html', 'xml']),
    enable_async=True
)

cache = {}

def get(name):
    if name not in cache:
        cache[name] = env.get_template(name)
    return cache[name]

async def render(name, request, **context):
    tmpl = get(name)
    return await tmpl.render_async(
        request=request,
        url_for=request.app.url_for,
        app=request.app,
        **context,
    )

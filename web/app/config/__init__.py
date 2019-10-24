"""
Handles configuration loading from environement variables and command line arguments
"""

import os

from envparse import Env
import click
from sanic import Sanic

env = Env(
    SERVER_NAME=dict(cast=str, default=""),
    HOST=dict(cast=str, default="127.0.0.1"),
    PORT=dict(cast=int, default=8080),
    DEBUG=dict(cast=bool, default=False),
)

def load(app: Sanic, *_):
    """Load the config from the environment"""
    env.read_envfile()
    app.config.SERVER_NAME = env('SERVER_NAME')
    app.config.HOST = env('HOST')
    app.config.PORT = env('PORT')
    app.config.DEBUG = env('DEBUG')

def run_params(f):
    """Decorator to add the config parameters to a Click command"""
    @click.option('--server-name', type=str, default=env('SERVER_NAME'))
    @click.option('-h', '--host', type=str, default=env('HOST'))
    @click.option('-p', '--port', type=int, default=env('PORT'))
    @click.option('--debug/--no-debug', default=env('DEBUG'))
    def new_func(server_name: str, host: str, port: int, debug: bool, *args, **kwargs):
        os.environ['HOST'] = host
        os.environ['PORT'] = str(port)
        os.environ['DEBUG'] = '1' if debug else '0'
        os.environ['SERVER_NAME'] = server_name
        return f(*args, **kwargs)

    return new_func

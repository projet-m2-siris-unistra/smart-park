"""
Handles configuration loading from environement variables
and command line arguments
"""

import os

import click
from envparse import Env
from sanic import Sanic

env = Env(
    SERVER_NAME=dict(cast=str, default=""),
    HOST=dict(cast=str, default="127.0.0.1"),
    PORT=dict(cast=int, default=8080),
    DEBUG=dict(cast=bool, default=False),
    NATS_URL=dict(cast=str, default="nats://localhost:4222"),
)


def load(app: Sanic, *_):
    """Load the config from the environment"""
    env.read_envfile()
    app.config.NATS_URL = env("NATS_URL")
    app.config.SERVER_NAME = env("SERVER_NAME")
    app.config.HOST = env("HOST")
    app.config.PORT = env("PORT")
    app.config.DEBUG = env("DEBUG")


def run_params(f):
    """Decorator to add the config parameters to a Click command"""

    @click.option("--nats-url", type=str, default=env("NATS_URL"))
    @click.option("--server-name", type=str, default=env("SERVER_NAME"))
    @click.option("-h", "--host", type=str, default=env("HOST"))
    @click.option("-p", "--port", type=int, default=env("PORT"))
    @click.option("--debug/--no-debug", default=env("DEBUG"))
    def new_func(
        nats_url: str,
        server_name: str,
        host: str,
        port: int,
        debug: bool,
        *args,
        **kwargs
    ):
        os.environ["HOST"] = host
        os.environ["PORT"] = str(port)
        os.environ["DEBUG"] = "1" if debug else "0"
        os.environ["SERVER_NAME"] = server_name
        os.environ["NATS_URL"] = nats_url
        return f(*args, **kwargs)

    return new_func

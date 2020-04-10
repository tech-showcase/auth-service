from flask import request
from functools import wraps
import env
from requests import get


def authentication_function(f):
    @wraps(f)
    def authenticate(*args, **kwargs):
        token = request.headers['Authorization']
        result = get(env.read_url_auth(), headers={'Authorization': token})

        if result.status_code != 200:
            return result.json(), result.status_code

        kwargs['claims'] = result.json()
        return f(*args, **kwargs)

    return authenticate


def admin_authorization_function(f):
    @wraps(f)
    def authorize_admin(*args, **kwargs):
        claims = kwargs['claims']
        if claims['role'] != "admin":
            return {'message': 'only admin is allowed to access'}, 401

        return f(*args, **kwargs)

    return authorize_admin

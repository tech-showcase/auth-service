from flask import Blueprint, jsonify, request
import controller
import env
from model import Resources
import data as data
from middleware import authentication_function, admin_authorization_function

api = Blueprint('api', __name__)


@api.route('/api/resources')
@authentication_function
def fetch_resources(*args, **kwargs):
    resources = Resources(env.read_url_resources())
    result_dict = controller.fetch_resources(resources, data.rupiah)
    return jsonify(result_dict)


@api.route('/api/resources/aggregate_price')
@authentication_function
@admin_authorization_function
def aggregate_resources(*args, **kwargs):
    date = request.args.get('date')
    province = request.args.get('province')

    resources = Resources(env.read_url_resources(), date_str=date, province=province)
    result_dict = controller.aggregate_resources_price(resources)
    return jsonify(result_dict)

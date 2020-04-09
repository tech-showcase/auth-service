from flask import Blueprint, jsonify
import controller
import env
from model import Resources
import data as data

api = Blueprint('api', __name__)


@api.route('/api/fetch_resources')
def fetch_resources():
    resources = Resources(env.read_url_resources())
    result_dict = controller.fetch_resources(resources, data.rupiah)
    return jsonify(result_dict)


@api.route('/api/aggregate_resources')
def aggregate_resources():
    resources = Resources(env.read_url_resources(), date_str='2020-02-02', province='JAWA TENGAH')
    result_dict = controller.aggregate_resources_price(resources)
    return jsonify(result_dict)

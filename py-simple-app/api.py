from flask import Blueprint, jsonify
import controller
from model import Resources
from data import rupiah
import env

api = Blueprint('api', __name__)


@api.route('/api/fetch_resources')
def fetch_resources():
    resources = Resources(env.read_url_resources())
    result_dict = controller.fetch_resources(resources, rupiah)
    return jsonify(result_dict)

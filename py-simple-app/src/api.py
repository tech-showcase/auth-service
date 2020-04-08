from flask import Blueprint, jsonify
from src import controller, env
from src.model import Resources
import src.data as data

api = Blueprint('api', __name__)


@api.route('/api/fetch_resources')
def fetch_resources():
    resources = Resources(env.read_url_resources())
    result_dict = controller.fetch_resources(resources, data.rupiah)
    return jsonify(result_dict)

from flask import Flask
from cmd import parse_args
from api import api
from data import update_rupiah_cache_every
from threading import Thread
import env

args = parse_args()

app = Flask(__name__)
app.register_blueprint(api)

if __name__ == '__main__':
    rupiah_checker_thread = Thread(target=update_rupiah_cache_every, args=(int(env.read_interval_rupiah_rate()),))
    rupiah_checker_thread.start()
    app.run(port=args.port)

from flask import Flask
from cmd import parse_args
from api import api

args = parse_args()

app = Flask(__name__)

if __name__ == '__main__':
    app.run(port=args.port)

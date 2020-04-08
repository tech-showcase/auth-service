from model import Rupiah
from time import sleep
import env

rupiah = None


def update_rupiah_cache_every(interval_sec):
    global rupiah
    url = env.read_url_rupiah_rate()
    rupiah = Rupiah(url)
    while True:
        rupiah.fetch()
        sleep(interval_sec)

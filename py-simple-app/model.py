from requests import get
from threading import Timer


class Resources:
    def __init__(self, url):
        self.url = url

    def fetch(self):
        result = get(self.url)
        result_dict = result.json()

        return result_dict


class Rupiah:
    def __init__(self, url, interval_sec):
        self.dollar_rate = 0.0
        self.url = url
        fetcher_thread = Timer(interval_sec, self.fetch)
        fetcher_thread.start()

    def fetch(self):
        result = get(self.url)
        result_dict = result.json()
        self.dollar_rate = result_dict['USD_IDR']

    def convert_to_dollar(self, price_in_rupiah):
        return self.dollar_rate * price_in_rupiah

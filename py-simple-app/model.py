from requests import get


class Resources:
    def __init__(self, url):
        self.url = url

    def fetch(self):
        result = get(self.url)
        result_dict = result.json()
        return result_dict


class Rupiah:
    def __init__(self, url):
        self.dollar_rate = 0.0
        self.url = url

    def fetch(self):
        result = get(self.url)
        result_dict = result.json()
        self.dollar_rate = result_dict['IDR_USD']
        print(self.dollar_rate, result_dict)

    def convert_to_dollar(self, price_in_rupiah):
        if price_in_rupiah is None:
            return None
        else:
            price_in_rupiah = float(price_in_rupiah)
            return self.dollar_rate * price_in_rupiah

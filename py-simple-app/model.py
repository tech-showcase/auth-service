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
        self.rate_against_dollar = 0.0
        self.url = url

    def fetch(self):
        result = get(self.url)
        result_dict = result.json()
        if result.status_code == 200:
            self.rate_against_dollar = result_dict['IDR_USD']
            print('fetched rupiah rate: ', self.rate_against_dollar)
        else:
            error = result_dict['error']
            print('failed to fetch rupiah rate: ', error)

    def convert_to_dollar(self, price_in_rupiah):
        if price_in_rupiah is None:
            return None
        else:
            price_in_rupiah = float(price_in_rupiah)
            return str(self.rate_against_dollar * price_in_rupiah)

from requests import get


class Resources:
    def __init__(self, url, price_converter):
        self.url = url
        self.price_converter = price_converter

    def fetch(self):
        result = get(self.url)
        result_dict = result.json()

        for index in range(len(result_dict)):
            price = result_dict[index]["price"]
            result_dict[index]["price_in_dollar"] = str(self.price_converter.convert_to_dollar(price))

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

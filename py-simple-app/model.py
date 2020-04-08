from requests import get


class Resources:
    def __init__(self, url):
        self.url = url

    def fetch(self):
        result = get(self.url)
        result_dict = result.json()
        return result_dict

from requests import get
from datetime import datetime, timedelta
import pandas as pd
import dateutil.parser


class Resources:
    def __init__(self, url, date_str=None, province=None):
        self.data = None
        self.price_statistics = None
        self.url = url

        if province is not None:
            self.province = province

        if date_str is not None:
            self.date_str = date_str
        self.start_of_week = None
        self.end_of_week = None

    def fetch(self):
        result = get(self.url)
        self.data = result.json()
        return self.data

    def aggregate_price(self):
        self.fetch()
        self.__get_start_date_on_same_week()
        self.__get_end_date_on_same_week()
        self.__filter_by_range_date_and_province()
        self.__get_price_statistics()
        return self.price_statistics

    def __filter_by_range_date_and_province(self):
        print(self.start_of_week, self.end_of_week)

        df = pd.DataFrame(self.data)
        df_non_null = df[df['tgl_parsed'].notna()]

        province_filtered_df = df_non_null.loc[df_non_null['area_provinsi'] == self.province]

        province_filtered_df['tgl'] = province_filtered_df['tgl_parsed'].apply(
            lambda tgl: dateutil.parser.isoparse(tgl).strftime('%Y-%m-%d'))

        date_filtered_df = province_filtered_df.loc[
            (province_filtered_df['tgl'] >= self.start_of_week.strftime('%Y-%m-%d')) &
            (province_filtered_df['tgl'] <= self.end_of_week.strftime('%Y-%m-%d'))
            ]

        self.data = date_filtered_df

    def __get_start_date_on_same_week(self):
        date_obj = datetime.strptime(self.date_str, '%Y-%m-%d')
        self.start_of_week = date_obj - timedelta(days=date_obj.weekday())

    def __get_end_date_on_same_week(self):
        self.end_of_week = self.start_of_week + timedelta(days=6)

    def __get_price_statistics(self):
        price_statistics = self.data['price'] \
            .dropna() \
            .apply(lambda price_str: int(price_str)) \
            .describe()
        print(price_statistics)
        self.price_statistics = dict()
        self.price_statistics['min'] = price_statistics['min']
        self.price_statistics['max'] = price_statistics['max']
        self.price_statistics['median'] = price_statistics['50%']
        self.price_statistics['avg'] = price_statistics['mean']


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

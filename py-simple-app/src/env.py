from os import environ


def read_url_rupiah_rate():
    var_name = 'URL_RUPIAH_RATE'
    return read(var_name)


def read_url_resources():
    var_name = 'URL_RESOURCES'
    return read(var_name)


def read_interval_rupiah_rate():
    var_name = 'INTERVAL_RUPIAH_RATE'
    return read(var_name)


def read(var_name):
    url = environ.get(var_name)
    if url is None:
        print('failed to read env var: ', var_name)

    return url

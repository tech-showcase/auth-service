def fetch_resources(resources, price_converter):
    result_dict = resources.fetch()

    for index in range(len(result_dict)):
        price = result_dict[index]["price"]
        result_dict[index]["price_in_dollar"] = price_converter.convert_to_dollar(price)

    return result_dict


def aggregate_resources_price(resources):
    result_dict = resources.aggregate_price()
    return result_dict

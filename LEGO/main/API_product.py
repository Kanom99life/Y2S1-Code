import requests
import json

def get_product(lang:str,sku:str)->dict:
    url = "https://ppe-api.lotuss.com/proc/product/api/v1/products/details?websiteCode=thailand_hy&storeId=5016&sku={}".format(sku)

    payload={}
    headers = {
    'accept-language': lang,
    'Authorization': 'Basic ZWY4ZTZjMjgzODdlNGVjYTlkM2UxMTU1MDQxMjgyYzE6MEU1NTg0QzUxZTdBNDBEODkzMDUxZGExY2NEQTg2ZTY='
    }

    response = requests.request("GET", url, headers=headers, data=payload)
    # print(response.text)
    response = json.loads(response.text)
    print(response.text)
    return response["data"]



def main():
    lang = 'th'
    sku = '51202528'
    print(get_product(lang,sku))
    
    
if __name__ == '__main__':
    main()
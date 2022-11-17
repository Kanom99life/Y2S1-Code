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
    return response["data"]


def cheangeTostring(data:dict)->str:
    result = ''
    result += str(data["name"]) + ","
    result += str(data["mediaGallery"][0]["url"]) + ","
    result += str(data["priceRange"]["minimumPrice"]["regularPrice"]["value"]) + " " + str(data["priceRange"]["minimumPrice"]["regularPrice"]["currency"]) + ","
    result += str(data["priceRange"]["minimumPrice"]["finalPrice"]["value"]) + " " + str(data["priceRange"]["minimumPrice"]["finalPrice"]["currency"]) + ","
    result += str(data["priceRange"]["minimumPrice"]["discount"]["percentOff"]) + ","
    if len(data["promotions"]) == 0:
        result += str("None")
    else:
        result += str(data["promotions"][0]["image"]) + "," + str(data["promotions"][0]["ruleType"])
    return result #name,image_product,regularPrice,finalPrice,percentOff,promotions,ruleType


def main():
    lang = 'th'
    sku = '51202528'
    P_DATA = get_product(lang,sku)
    cheangeTostring(P_DATA)
    print(cheangeTostring(P_DATA))
    
if __name__ == '__main__':
    main()
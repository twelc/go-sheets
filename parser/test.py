import requests
import bs4

q = "https://krisha.kz/arenda/kvartiry/almaty/"

headers = {
    "Host": "krisha.kz",
    "User-Agent": "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/119.0",
    "Accept": "application/json, text/javascript, */*; q=0.01",
    "Accept-Language": "ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3",
    "Accept-Encoding": "gzip, deflate, br",
    "X-Requested-With": "XMLHttpRequest",
    "Alt-Used": "krisha.kz",
    "Connection": "keep-alive",
    "Referer": "https://krisha.kz/arenda/kvartiry/almaty/?das[map.complex]",
    "Sec-Fetch-Dest": "empty",
    "Sec-Fetch-Mode": "cors",
    "Sec-Fetch-Site": "same-origin",
    "TE": "trailers"
    }

r = requests.get(q, headers=headers)

soup = bs4.BeautifulSoup(r.text, "html.parser")
print(soup)
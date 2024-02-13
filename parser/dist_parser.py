import requests
from bs4 import BeautifulSoup as bs4
import json
from fuzzywuzzy import process
import socks
import socket

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

def _p(r):
    soup = bs4(r.text, "html.parser")
    all = soup.findAll("div", class_="complex-card__child-block")
    res = {}
    for item in all:
        comp = item.findChild("p", class_="complex-card__title")
        dist = item.findChild("p", class_="complex-card__address")
        print(comp.text.lower().strip())
        print(dist.text.lower().strip())
        print(dist.text.lower().split(",")[1].strip())
        print("\033[94m========\033[0m")
        res[comp.text.lower().strip()] = dist.text.lower().split(",")[1].strip()
    return res
    
def checkIP():
        ip = requests.get('https://api.ipify.org/?format=json')
        print('ip is: ' + '\033[92m'+ ip.json()["ip"] + '\033[0m')

def parse():
    socks.set_default_proxy(socks.SOCKS5, "localhost", 9150)
    socket.socket = socks.socksocket
    res = {}
    q = "https://krisha.kz/complex/search/almaty/?state[]="
    q2 = "https://krisha.kz/complex/search/almaty/?state[]=&page={}"
    r = requests.get(url=q, headers=headers)
    
    _res = _p(r)
    
        
    res = res | _res
    for i in range(1, 56):
        print("\033[94m=====================================\033[0m")
        checkIP()
        r = requests.get(url=q2.format(str(i)), headers=headers)
        try:
            _res = _p(r)
        except TypeError:
            print(f"\033[93mpage {str(i)}\033[0m")
        res = res | _res
    with open("d2.json", "w+") as file:
        json.dump(res, file, ensure_ascii=False, indent=4)


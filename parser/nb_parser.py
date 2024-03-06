import requests
import districts
import datetime
import socks
import socket
import writer
import sheet_lib
import time



from loguru import logger
print("started krisha parsing")
dists = districts.district_manager()
querry = "https://krisha.kz/a/ajaxGetSearchNbResults?search-url=/arenda/kvartiry/almaty/?das[map.complex]={}&isOnMap=0"
querry2 = "https://krisha.kz/arenda/kvartiry/almaty/?das[map.complex]={}"
sheet_lib.make_config('./roofsparser-addef44f7a5a.json',
                      "141maOrpeeFsydVAWP-kIaziMCHn_fI8nQv0mFB78TVk",
                      "history")

Global_index = 0

logger.add("log.log")

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

def checkIP():
        hostname = socket.gethostname()
        ip_address = socket.gethostbyname(hostname)
        print('ip is: ' + '\033[92m'+ ip_address + '\033[0m')


def write(data:str):
    with open("result.txt", "+a") as file:
        file.write(data + "\n")



def parse_data():
    i=0
    index = 1
    writer.create_tmp()
    socks.set_default_proxy(socks.SOCKS5, "localhost", 9150)
    socket.socket = socks.socksocket
    all_start = datetime.datetime.now()
    while True:
        try:
            start = datetime.datetime.now()
            i+=1
            if i > 3500:
                break

            r = requests.get(url=querry.format(str(i)), headers=headers)
            values = ('"nb":', '}')  
            nb = r.text[r.text.find(values[0]) + len(values[0]):]
            nb = nb[:nb.find(values[1])]
            r = requests.get(url=querry2.format(str(i)), headers=headers)
            
            txt = r.text

            
            values = ('Аренда квартир помесячно в ЖК', 'в Алматы')  
            csrf = txt[txt.find(values[0]) + len(values[0]):]
            csrf = csrf[:csrf.find(values[1])]
            d = dists.get_dist(csrf.lower().strip())
            try:
                int(csrf.lower().strip())
            except:
                write(f"{d}, {csrf}: {str(nb)}")
                writer.save_data([[d, csrf, str(nb)]], index)
                sheet_lib.append_data(csrf, d, str(nb), Global_index)
                index += 1
            else:
                continue
        except KeyboardInterrupt:
            raise KeyboardInterrupt
        except:
            logger.exception("error")
    writer.accept_dump()
    

    
    

if __name__ == "__main__":
    next = datetime.datetime.now() + datetime.timedelta(days=1)
    parse_data()
    while True:
        time.sleep(10)
        if datetime.datetime.now() > next:
            Global_index+=1
            next = datetime.datetime.now() + datetime.timedelta(days=1)
            parse_data()


import requests
from bs4 import BeautifulSoup as bs4
import json
from fuzzywuzzy import process



class district_manager:
    def __init__(self) -> None:
        self.districts = self.get_fdistricts()
        print(self.districts)


    def get_fdistricts(self) -> dict:
        with open("d2.json", "r") as file:
            dists = json.load(file)
            return dists

    def get_districts(self) -> dict:
        res = {}
        q = "https://homsters.kz/estate/search?TypeOfSearchComplexOrEstate=1&AdministrativeUnitType=City&AdminUnitsId=993&page={}"
        page = 0
        while True:
            page+=1
            print("page: " + str(page))
            if page>30:
                break
            r = requests.get(url=q.format(str(page)))
            soup = bs4(r.text, "html.parser")
            wrapper = soup.find("div", class_="b-search__list js-search-list js-search-complex-list style-grid active")
            try:
                items = wrapper.findChildren("div", class_="b-snippet__wrapper swiper-slide js-project-item")
            except AttributeError:
                break
            for item in items:
                rc = str(item.findChild("h2").text).lower().replace("жк", "").strip()
                dist = str(item.findChild("div", class_="b-snippet__location-text").text).lower().strip()
                res[rc] = dist
                print(rc + ": " + dist)
        with open("d2.json", "+w") as file:
            json.dump(res, file)
        self.districts = res
        return res
    
    def get_dist(self, dist):
        a = process.extractOne(dist, self.districts.keys())
        print(a)
        if a[1] <= 80:
            return "Район не определён"
        return self.districts[a[0]]


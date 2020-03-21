import requests
import re
from bs4 import BeautifulSoup

url = 'https://www.yahoo.co.jp/'
r = requests.get(url)

soup = BeautifulSoup(r.content, 'html.parser')

arefs = soup.find_all(href=re.compile("news.yahoo.co.jp/pickup"))
for i, aref in enumerate(arefs):
    print(aref.span.string)
    print(aref.attrs['href'])
    if i != len(arefs) - 1:
        print('\n')
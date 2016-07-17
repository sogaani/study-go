#!/usr/bin/python
import pyquery

for page in range(20):
    query = pyquery.PyQuery("http://www.alexa.com/topsites/global;" + str(page) , parser='html')
    for li in query('.site-listing')('li'):
        print query(li)('.count').text() + ", " + query(li)('.desc-paragraph')('a').text()

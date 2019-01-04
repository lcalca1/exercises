#!/usr/local/bin/python3
# coding=utf-8

import time
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.chrome.options import Options

options = Options()
options.add_argument('--headless')
options.add_argument('--no-sandbox')


driver = webdriver.Chrome('/bin/chromedriver', chrome_options=options)
driver.get("http://player.youku.com/embed/XMzkyNTkwMjc2MA==")
#driver.get("https://www.youku.com/")
time.sleep(55)
'''
assert "Python" in driver.title
print driver.title
elem = driver.find_element_by_name("q")
elem.clear()
elem.send_keys("pycon")
elem.send_keys(Keys.RETURN)
assert "No results found." not in driver.page_source
'''
#print( driver.execute_script('return window[ window.UA_Opt.LogVal ]') )
print( driver.execute_script('return window._collina.ckey') )
driver.close()

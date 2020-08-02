from selenium import webdriver

driver = webdriver.Firefox(executable_path='/usr/local/bin/geckodriver')
driver.get('https://www.google.co.jp')
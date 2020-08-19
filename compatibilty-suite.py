import sys
import os
import signal
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.common.exceptions import TimeoutException
from selenium.webdriver.support import expected_conditions as EC
from subprocess import Popen


# launch node

p = None

if len(sys.argv) > 1:
   if sys.argv[1] == "record":
       url = "http://localhost:61417"
       p = Popen(["go", "run", "todobackend_compatibility.go", "record"])
   elif sys.argv[1] == "playback":
       url = "http://localhost:61417"
       p = Popen(["go", "run", "todobackend_compatibility.go", "playback"])
   elif sys.argv[1] == "direct":
       print("showing reference Sinatra app online without Servirtium in the middle")
       url = "https://todo-backend-sinatra.herokuapp.com"
   else:
       print("Second arg should be record or playback")
       exit(10)
else:
   print("record/playback/direct argument needed")
   exit(10)

driver = webdriver.Chrome("/usr/local/bin/chromedriver")

driver.get("https://www.todobackend.com/specs/index.html?" + url + "/todos")
try:
    element = WebDriverWait(driver, 300).until(
        EC.text_to_be_present_in_element((By.CLASS_NAME, "passes"), "16")
    )
    print("Compatibility suite: all 16 tests passed")
except TimeoutException as ex:
    print("Compatibility suite: did not finish with 16 passes. See open browser frame.")
if p is not None:
    os.killpg(os.getpgid(p.pid), signal.SIGTERM)
    p.kill()
driver.quit()
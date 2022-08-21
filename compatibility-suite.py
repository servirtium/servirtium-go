import sys
import os
import signal
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.common.exceptions import TimeoutException
from selenium.webdriver.support import expected_conditions as EC
import subprocess
import time

servirtium_process = None
todoSuiteUrl = "https://servirtium.github.io/compatibility-suite/index.html"

if len(sys.argv) > 1:

   realUrl = "https://http4k-todo-backend.herokuapp.com/"

   if sys.argv[1] == "record":
       # TODO check that node process is already started.
       url = "http://localhost:61417"
       servirtium_process = subprocess.Popen(["go", "run", "./cmd/todobackend_compatibility.go", "record", realUrl])
   elif sys.argv[1] == "playback":
       url = "http://localhost:61417"
       servirtium_process = subprocess.Popen(["go", "run", "./cmd/todobackend_compatibility.go", "playback", realUrl])
   elif sys.argv[1] == "direct":
       print("showing Http4k Todobackend implementation online without Servirtium in the middle")
       todoSuiteUrl = "https://www.todobackend.com/specs/index.html"
       url = realUrl
   else:
       print("Second arg should be record or playback")
       exit(10)
else:
   print("record/playback/direct argument needed")
   exit(10)

if len(sys.argv) > 2:
   if sys.argv[2] == "local":
       # Running via 'python -m SimpleHTTPServer 8000' in a different shell
       todoSuiteUrl = "http://localhost:8000/index.html"

driver = webdriver.Chrome("chromedriver")

time.sleep(5) # for old workstations

driver.get(todoSuiteUrl + "?" + url + "/todos")
try:
    element = WebDriverWait(driver, 300).until(
        EC.text_to_be_present_in_element((By.CLASS_NAME, "passes"), "16")
    )
    print("Compatibility suite: all 16 tests passed")

except TimeoutException as ex:
    print("Compatibility suite: did not finish with 16 passes. See open browser frame.")

# TODO warn that node process was not started.

print("mode: " + sys.argv[1])

if servirtium_process is not None:
    print("Killing servirtium process")
    os.killpg(os.getpgid(servirtium_process.pid), signal.SIGTERM)
    servirtium_process.kill()

print("Closing Selenium")
driver.quit()
print("All done.")

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
import docker
import signal

servirtium_process = None
dkrCtr = None


# Pete Hodgson's original is at https://www.todobackend.com/specs/index.html
# Ours has a few additions to aid its use in this compatibility suite
# Playback and Record mode use ours. Direct mode uses Pete's original.
todoSuiteUrl = "https://servirtium.github.io/compatibility-suite/index.html"


# We used to use the Ruby-Sinatra version of the Todobackend, but it is not
# up as often as the http4k version, but both are unavailable as Heroku changed
# their free model.
# realUrl = "https://todo-backend-sinatra.herokuapp.com"
realUrl = "https://todo-backend-rocket-rust.herokuapp.com/"
# So we use our own version of http4k-todo-backend via Docker.
#  See https://github.com/servirtium/todobackend-for-compatibility-kit for _build_ instructions
# realUrl = "http://localhost:54321"
extraURL = ""

docker_stopped = False
servirtium_stopped = False
def stop_docker():
    global docker_stopped
    if docker_stopped is not True and dkrCtr is not None:
        print("Killing Locally launched Todo-backend")
        dkrCtr.stop()
        docker_stopped = True

def stop_servirtium():
    global servirtium_stopped

    if servirtium_stopped is not True and servirtium_process is not None:
        print("=== Servirtium process === ")
        for line in iter(servirtium_process.stdout.readline, b''):
            sys.stdout.write(line.decode(sys.stdout.encoding))
        print("=== Killing servirtium process ===")
        try:
            os.killpg(os.getpgid(servirtium_process.pid), signal.SIGTERM)
        except ProcessLookupError:
            pass
        servirtium_process.kill()
        servirtium_stopped = True

def ctrlc_handler(signum, frame):
    stop_docker()
    stop_servirtium()

signal.signal(signal.SIGINT, ctrlc_handler)

if len(sys.argv) > 1:

    try:
        if "54321" in realUrl and (sys.argv[1] == "record" or sys.argv[1] == "direct"):
            dkr = docker.from_env()
            dkrCtr = dkr.containers.run("todobackend-api-for-servirtium-development", ports={'54321/tcp': 54321}, detach=True)
    except BaseException as err:
        if "port is already allocated" in str(err):
            print("docker ps - then kill the todobackend-api-for-servirtium-development already running")
            exit(10)

    if sys.argv[1] == "record":
        # TODO check that node process is already started.
        url = "http://localhost:61417"
        servirtium_process = subprocess.Popen(["go", "run", "./cmd/todobackend_compatibility.go", "record", realUrl], stderr=subprocess.STDOUT, stdout=subprocess.PIPE)
    elif sys.argv[1] == "playback":
        url = "http://localhost:61417"
        servirtium_process = subprocess.Popen(["go", "run", "./cmd/todobackend_compatibility.go", "playback", realUrl], stderr=subprocess.STDOUT, stdout=subprocess.PIPE)
    elif sys.argv[1] == "direct":
        print("showing Http4k Todobackend implementation online without Servirtium in the middle")
        extraURL = "&noServirtium"
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

print("sleep 5 seconds for older workstations")
time.sleep(5)

driver.get(todoSuiteUrl + "?" + url + extraURL)
try:
    element = WebDriverWait(driver, 300).until(
        EC.text_to_be_present_in_element((By.CLASS_NAME, "passes"), "16")
    )
    print("Compatibility suite: all 16 tests passed")
except BaseException as ex:
    print("Compatibility suite: DID NOT finish with 16 passes. See browser frame.")

print("mode: " + sys.argv[1])

stop_servirtium()

stop_docker()

print("Closing Selenium")
driver.quit()
print("All done.")

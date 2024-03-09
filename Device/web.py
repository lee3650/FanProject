import time
import requests

server_url = "temp"
stored_data = None
running = False
device_id = 0


def initialize():
    running = True
    get_device_id()
    query_server()


def get_device_id():
    return  # TODO


def update(new_data):
    return  # TODO


def query_server():
    while running:
        http_response = requests.get(f"{server_url}/readstate?id={device_id}")
        data = http_response.json()
        if data != stored_data:
            update(data)
        time.sleep(1)





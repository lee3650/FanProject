from flask import Flask, request
import os
import subprocess
import requests

device = Flask(__name__)

@device.route("/", methods=["GET"])
def hello_world():
    return "<p>Hello, World!</p>"

@device.route("/connect", methods=["POST"])
def connect_to_internet():
    data = request.get_json()
    # change this to raspberry pi command
    output = subprocess.check_output(f"networksetup -setairportnetwork en0 '{data['ssid']}' '{data['password']}'", shell=True).decode()
    if "Error: -" in output:
        return {"message": output}, 400
    
    try:
        # change this to central server address
        requests.get("http://google.com")
    except:
        return {"message": "connected to LAN but not to server"}, 500
    
    return {"message": "connected to the internet"}
    

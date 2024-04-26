from flask import Flask, jsonify, request, render_template
import requests

app = Flask(__name__)

# get data of an user
url = "http://localhost:8080/user"
r = requests.get(url = url)
data = r.json

@app.route('/returnjson', methods = ['GET']) 
def ReturnJSON():
    return data 

@app.route('/', methods = ['GET'])
def home():
    return render_template("home.html", title="CV Project")

@app.route('/template1', methods = ['GET'])
def generate_template():
    return render_template("template1.html")

  
if __name__=='__main__': 
    app.run(debug=True)
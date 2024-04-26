from flask import Flask, render_template
import requests

app = Flask(__name__)

@app.route('/', methods = ['GET'])
def home():
    return render_template("home.html", title="CV Project")

@app.route('/template1', methods = ['GET'])
def generate_template():
    url = "http://localhost:8080/user"
    r = requests.get(url = url)
    data = r.json()
    return render_template("template1.html", data = data)

  
if __name__=='__main__': 
    app.run(debug=True)
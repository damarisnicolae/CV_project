from flask import Flask, render_template
import requests

app = Flask(__name__)

@app.route('/', methods=['GET'])
def home():
    url = "http://localhost:8080/home"

    u = requests.get(url = url)
    data = u.json()
    return render_template('home.html', data = data)

@app.route('/template1', methods = ['GET'])
def generate_template1():
    url = "http://localhost:8080/user"
    url2 = "http://localhost:8080/pdf?template=1&user=1" ####todo
    
    r = requests.get(url = url)
    data = r.json()
    
    r2 = requests.get(url = url2) ####todo
    
    return render_template("template1.html", data = data)

@app.route('/template2', methods = ['GET'])
def generate_template2():
    url = "http://localhost:8080/user"
    url2 = "http://localhost:8080/pdf"
    
    r = requests.get(url = url)
    data = r.json()
    
    return render_template("template2.html", data = data)

@app.route('/template3', methods = ['GET'])
def generate_template3():
    url = "http://localhost:8080/user"
    url2 = "http://localhost:8080/pdf"
    
    r = requests.get(url = url)
    data = r.json()
    
    return render_template("template3.html", data = data)

if __name__=='__main__': 
    app.run(debug=True)
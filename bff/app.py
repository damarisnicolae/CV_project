from flask import Flask, render_template, request
import requests, json

app = Flask(__name__)

@app.route('/', methods=['GET'])
def home():
    url = "http://localhost:8080/home"

    u = requests.get(url = url)
    data = u.json()
    return render_template('home.html', users = data)

@app.route('/user', methods = ['POST'])
def add_user():
    url = "http://localhost:8080/user"
    requests.post(url = url)
    return render_template('home.html')

@app.route('/user', methods = ['GET'])
def get_user():
    url = "http://localhost:8080/user"
    u = requests.get(url = url)
    data = u.json()
    return render_template('edit_form.html', data = data)

@app.route('/user', methods = ['PUT'])
def edit_user():
    edited_data = request.json
    
    url = "http://localhost:8080/user/1"
    headers = {"Content-Type": "application/json"}
    requests.put(url, data=json.dumps(edited_data), headers=headers)
    
    return render_template('home.html')

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

@app.route('/userlogin', methods = ['GET','POST'])
def loginuser():
    if request.method == "GET":
        return render_template('loginform.html')
    if request.method == "POST":
        r = requests.post(f'http://localhost:8080/user', request.form, headers=request.headers)
        response_json = r.json()

if __name__=='__main__': 
    app.run(debug=True)
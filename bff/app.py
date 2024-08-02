from flask import Flask, render_template, request
import requests, json, argparse

parser = argparse.ArgumentParser()
parser.add_argument("-i", "--ip", help="API IP")
parser.add_argument("-p", "--port", help="API PORT")
args = vars(parser.parse_args())

IP = args["ip"]
PORT = args["port"]

app = Flask(__name__)

@app.route('/', methods=['GET'])
def home():
    url = f"http://{IP}:{PORT}/users"
    response = requests.get(url = url)
    data = response.json()
    return render_template('home.html', users = data)

@app.route('/user', methods = ['POST'])
def add_user():
    url = f"http://{IP}:{PORT}/user"
    requests.post(url = url)
    return render_template('home.html')

@app.route('/user/<id>', methods = ['PUT'])
def edit_user(id):
    edited_data = request.json
    url = f"http://{IP}:{PORT}/user/{id}"
    headers = {"Content-Type": "application/json"}
    requests.put(url, data=json.dumps(edited_data), headers=headers)
    return render_template('home.html')

@app.route('/postform', methods = ['GET'])
def get_postform():
    return render_template('post_form.html')

@app.route('/editform', methods = ['GET'])
def get_user():
    url = f"http://{IP}:{PORT}/user"
    u = requests.get(url = url)
    data = u.json()
    return render_template('edit_form.html', data = data)

@app.route('/template1', methods = ['GET'])
def generate_template1():
    url = f"http://{IP}:{PORT}/user"
    url2 = f"http://{IP}:{PORT}/pdf?template=1&user=1" ####todo
    
    r = requests.get(url = url)
    data = r.json()
    
    r2 = requests.get(url = url2) ####todo
    
    return render_template("template1.html", data = data)

@app.route('/template2', methods = ['GET'])
def generate_template2():
    url = f"http://{IP}:{PORT}/user"
    url2 = f"http://{IP}:{PORT}/pdf"
    
    r = requests.get(url = url)
    data = r.json()
    
    return render_template("template2.html", data = data)

@app.route('/template3', methods = ['GET'])
def generate_template3():
    url = f"http://{IP}:{PORT}/user"
    url2 = f"http://{IP}:{PORT}/pdf"
    
    r = requests.get(url = url)
    data = r.json()
    
    return render_template("template3.html", data = data)

@app.route('/login', methods = ['GET','POST'])
def loginuser():
    if request.method == "GET":
        return render_template('loginform.html')
    if request.method == "POST":
        r = requests.post(f"http://{IP}:{PORT}/login", request.form, headers=request.headers)
        if r.status_code == 200:
            data = r.json()
            return render_template('greet.html', data = data)
        else:
            return render_template('loginform.html')
    
@app.route('/logout', methods = ['POST'])
def logoutuser():
    r = requests.post(f"http://{IP}:{PORT}/logout")
    return render_template('home.html')

@app.route('/signup', methods = ['GET','POST'])
def signupuser():
    if request.method == "GET":
        return render_template('signupform.html')
    if request.method == "POST":
        r = requests.post(f'http://{IP}:{PORT}/signup', request.form, headers=request.headers)
        data = r.json()
        return render_template('home.html', data = data)
    
if __name__=='__main__': 
    app.run(debug=True)

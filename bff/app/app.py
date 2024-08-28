from flask import Flask, render_template, request
import requests, json, argparse, os

print("Start")

parser = argparse.ArgumentParser()
parser.add_argument("-i", "--ip", help="API IP", default="cv_api_container")
parser.add_argument("-p", "--port", help="API PORT", default="8080")
args = vars(parser.parse_args())

IP = args["ip"]
PORT = args["port"]

# value = os.getenv('API_IP')

app = Flask(__name__, template_folder='../templates')


@app.route('/', methods=['GET'])
def home():
    print("Sunt in /")
    url = f"http://{IP}:{PORT}/users"
    try:
        response = requests.get(url=url)
        response.raise_for_status()  # Raise an HTTPError for bad responses
        data = response.json()
    except requests.exceptions.RequestException as e:
        app.logger.error(f"Request failed: {e}")
        abort(500, description="Internal Server Error")
    return render_template('view/home.html', users=data)

@app.route('/user', methods = ['POST'])
def add_user():
    url = f"http://{IP}:{PORT}/user"
    requests.post(url = url)
    return render_template('view/home.html')

@app.route('/user/<id>', methods = ['PUT'])
def edit_user(id):
    edited_data = request.json
    url = f"http://{IP}:{PORT}/user/{id}"
    headers = {"Content-Type": "application/json"}
    requests.put(url, data=json.dumps(edited_data), headers=headers)
    return render_template('view/home.html')

@app.route('/postform', methods = ['GET'])
def get_postform():
    return render_template('forms/post_form.html')

@app.route('/editform', methods = ['GET'])
def get_user():
    url = f"http://{IP}:{PORT}/user"
    u = requests.get(url = url)
    data = u.json()
    return render_template('forms/edit_form.html', data = data)

@app.route('/template1', methods = ['GET'])
def generate_template1():
    url = f"http://{IP}:{PORT}/user"
    url2 = f"http://{IP}:{PORT}/pdf?template=1&user=1" ####todo
    
    r = requests.get(url = url)
    data = r.json()
    
    r2 = requests.get(url = url2) ####todo
    
    return render_template("view/template1.html", data = data)

@app.route('/template2', methods = ['GET'])
def generate_template2():
    url = f"http://{IP}:{PORT}/user"
    url2 = f"http://{IP}:{PORT}/pdf"
    
    r = requests.get(url = url)
    data = r.json()
    
    return render_template("view/template2.html", data = data)

@app.route('/template3', methods = ['GET'])
def generate_template3():
    url = f"http://{IP}:{PORT}/user"
    url2 = f"http://{IP}:{PORT}/pdf"
    
    r = requests.get(url = url)
    data = r.json()
    
    return render_template("view/template3.html", data = data)

@app.route('/login', methods = ['GET','POST'])
def loginuser():
    if request.method == "GET":
        return render_template('forms/loginform.html')
    if request.method == "POST":
        r = requests.post(f"http://{IP}:{PORT}/login", request.form, headers=request.headers)
        if r.status_code == 200:
            # data = r.json()
            return render_template('view/greet.html')
        else:
            return render_template('forms/loginform.html')
    
@app.route('/logout', methods = ['GET'])
def logoutuser():
    r = requests.post(f"http://{IP}:{PORT}/logout")
    return render_template('view/home.html')

@app.route('/signup', methods = ['GET','POST'])
def signupuser():
    if request.method == "GET":
        return render_template('forms/signupform.html')
    if request.method == "POST":
        r = requests.post(f'http://{IP}:{PORT}/signup', request.form, headers=request.headers)
        return render_template('view/home.html')
    
if __name__=='__main__': 
    app.run(host='0.0.0.0', port=5000, debug=True)

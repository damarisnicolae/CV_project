from flask import Flask, abort, jsonify, render_template, send_from_directory, request
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
    
@app.route('/user', methods=['DELETE'])
def delete_user():
    user_id = request.args.get('id')

    if not user_id:
        app.logger.error("user id is missing")
        return jsonify({'error': 'user id required'}), 400
    
    url = f"http://{IP}:{PORT}/user/{user_id}"
    app.logger.info(f"Sending DELETE request to {url}")
    try:
        response = requests.delete(url)
        response.raise_for_status()  # Raise an HTTPError for bad responses
        app.logger.info(f"Successfully deleted user with ID {id}")
    except requests.exceptions.HTTPError as http_err:
        app.logger.error(f"HTTP error occurred: {http_err}")
        return jsonify({'error': 'User deletion failed', 'details': str(http_err)}), 500
    except requests.exceptions.RequestException as req_err:
        app.logger.error(f"Request error occurred: {req_err}")
        return jsonify({'error': 'User deletion failed', 'details': str(req_err)}), 500
    return jsonify({'message': f'User {user_id} delete successfully'}), 200

@app.route('/styles/<path:filename>')
def serve_css(filename):
    return send_from_directory(os.path.join(app.root_path, '../static/styles'), filename)

@app.route('/js/<path:filename>')
def serve_js(filename):
    return send_from_directory(os.path.join(app.root_path, '../static/js'), filename)

if __name__=='__main__': 
    app.run(host='0.0.0.0', port=5000, debug=True)
    

from flask import Flask, render_template

app = Flask(__name__)

users = [
    {"id": 1, "firstname": "Sherlock", "lastname": "Holmes", "email": "sholmes@example.com"},
    {"id": 2, "firstname": "Nicolae", "lastname": "Ceausescu", "email": "ceausescu_ro@example.com"}
]

@app.route('/', methods=['GET'])
def home():
    return render_template('home.html', users=users)

if __name__ == '__main__':
    app.run(debug=True)
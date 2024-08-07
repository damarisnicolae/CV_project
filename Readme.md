# CV Management System API

## Overview : _This project provides a platform for managing user data and generating custom CV templates._

```sh
CV_project
├── api                                 # Go API         - backend
│   ├── CV_project                      # Executable file ...  (cd $PathCvProject/api && go build -o CV_project main.go)
│   ├── example.pdf                     # Sample PDF
│   ├── go.mod                          # Go modules     - dependency management
│   ├── go.sum                          # Checksum       - dependencies
│   └── main.go                         # Main entry point
├── bff                                 # Flask-based    - frontend
│   ├── app.py                          # Main Flask application file
│   ├── css                             # CSS files      - styling
│   │   └── users.css                   # Custom CSS     - user-related pages
│   ├── templates                       # HTML templates - frontend
│   │   ├── edit_form.html              # HTML     form  - editing user data
│   │   ├── greet.html                  # Greeting page
│   │   ├── home.html                   # Home     page
│   │   ├── loginform.html              # Login    form
│   │   ├── populate_template.html      # Template       - populating CVs
│   │   ├── post_form.html              # Form           - posting new content
│   │   ├── signupform.html             # Signup form
│   │   ├── template1.html              # CV template 1
│   │   ├── template2.html              # CV template 2
│   │   └── template3.html              # CV template 3
│   └── users.js                        # JavaScript     - user-related functionality
├── .gitignore                          # Git ignore     - version control
├── sql                                 # SQL files      - database schema
│   ├── schemadump.sql                  # Schema creation and sample data
│   └── schema.sql                      # Schema creation only
└── src                                 # Source code
    └── __init__.py                     # Initialization and configuration
```

## Components:

    Backend  (Go)    : Handles user data management, authentication, and PDF generation.
    Frontend (Flask) : Provides a web interface for user interactions, including data input, template selection, and CV generation.

## Prerequisites:

- `Go                        `: _Backend development_
- `Flask                     `: _Frontend development_
- `MySQL database            `: _Storing user data and templates_
- `wkhtmltopdf               `: _PDF generation_
- `Git                       `: _Version control_
- `Docker & Docker Compose   `: _Containerized deployment_

## Install Basic Tools:

```sh
sudo apt update && sudo apt upgrade && sudo apt install -y git curl build-essential golang-go python3 python3-pip wkhtmltopdf docker.io docker-compose selinux-utils curl mysql-server
sudo mysql_secure_installation
pip install --break-system-packages Flask Flask-Bcrypt Flask-Migrate Flask-SQLAlchemy
```

## Replace Paths

```sh
PathCvProject="/bcn/github/CV_project"
grep -q "PathCvProject=" ~/.bashrc || echo "export PathCvProject=\"$PathCvProject\"                                         # Set path to CV project." >> ~/.bashrc && source ~/.bashrc
```

## Create DB, user

```sh
sudo mysql -u root -p
CREATE DATABASE IF NOT EXISTS users;
CREATE USER 'cv_user'@'localhost' IDENTIFIED BY 'Y0ur_strong_password';
GRANT ALL PRIVILEGES ON users.* TO 'cv_user'@'localhost';
FLUSH PRIVILEGES;
sudo mysql -u 'cv_user' -p users

```

## Change DB user password

```sh
ALTER USER 'cv_user'@'localhost' IDENTIFIED BY 'Y0ur_strong_password';
```

## Update DB temporary, import schemas

```sh
mysql -u 'cv_user' -p -e "USE cv_project; DROP TABLE IF EXISTS user, template;"
mysql -u 'cv_user' -p cv_project < $PathCvProject/sql/schemadump.sql  # with    user
mysql -u 'cv_user' -p cv_project < $PathCvProject/sql/schema.sql      # without user
```

## Verify successful import

```sh
mysql -u cv_user -p cv_project
SHOW DATABASES;
SHOW TABLES;
DESCRIBE user;
DESCRIBE template;
```

## Build the backend API

```sh
cd $PathCvProject/api
go build -o CV_project main.go
export DBUSER="cv_user"
export DBPASS="Y0ur_strong_password"
echo $DB_USER
echo $DB_PASSWORD
./CV_project
```

## BFF Flask app setup frontend

```sh
cd $PathCvProject/bff
export FLASK_APP=app
export FLASK_ENV=development
python3 app.py -i 127.0.0.1 -p 5000
```

# Github

### SSH conection

```sh
GitSshKey="/PathTo/.ssh/github_rsa"
GitUsername="YourUsername"
GitEmail="YourEmail"
chmod 600 "$GitSshKey"
ssh-add "$GitSshKey"
git config --global user.name "$GitUsername"
git config --global user.email "$GitEmail"
git config --global http.sslBackend "openssl"
ssh -T git@github.com
```

### Commit & pull-push, avoid conflicts

```sh
echo "Enter commit message (Title case, infinitive verb, brief and clear summary of changes):"
read -p "CommitMssg: - " CommitMssg
cd "$PathCvProject" || exit
git add .
git commit -m "$CommitMssg"
git pull && git push origin main
```

## Start the project

```sh
cd $PathCvProject && make
```

## Browser

http://127.0.0.1:5000/template1

http://127.0.0.1:5000/template2

http://127.0.0.1:5000/template3

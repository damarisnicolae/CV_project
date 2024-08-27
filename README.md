### **CV Management System API  📜 ✍️**
<br>

**This project provides a platform for managing user data and generating custom CV templates.**  
<br>

![](https://i.imgur.com/waxVImv.png)
<br>

![freeCodeCamp Social Banner](https://venngage-wordpress.s3.amazonaws.com/uploads/2021/11/section-3-resume-banner-1-1.png)
<div align="center" markdown="1">

![](https://i.imgur.com/waxVImv.png)

<br>

[![GitHub Repo stars](https://img.shields.io/badge/Project%20mindmap-Miro-yellow?logo=miro)](https://miro.com/app/board/uXjVK6HA_1A=/)
[![GitHub Repo stars](https://img.shields.io/badge/Thanks%20Contributors-❤️-red?logo=github)](https://github.com/damarisnicolae/CV_project/graphs/contributors/)
</div>
 <!-- <p align="center"><img src="https://cdn6.aptoide.com/imgs/9/2/6/9262d372de5cd29430c675a6099e115c_icon.png?w=128" height="128" ></p>  -->


<br>

### **Table of contents**
<br>

1.  [Directory structure Hierarchy](#directory-structure-hierarchy)
2.  [Components](#components)
3.  [Prerequisites](#prerequisites)
4.  [Environment variables file .env file](#environment-variables-file-env-file)
5.  [Install Basic Tools](#install-basic-tools)
6.  [Replace Paths](#replace-paths)
7.  [Create DB, users table](#create-db-users-table)
8.  [Verify successful import](#verify-successful-import)
9.  [Change DB user password](#change-db-user-password)
10. [Build the backend API](#build-the-backend-api)
11. [BFF Flask app setup frontend](#bff-flask-app-setup-frontend)
12. [SSH conection](#github-ssh-conection)
13. [Github commit](#github-commit)
14. [Start the project](#start-the-project-makefile)
15. [Docker](#docker)
16. [Browser links](#browser-links)
17. [Other free resources](#other-free-resources)  
<br>

### **Directory structure Hierarchy**
<br>


```sh
CV_project
├── .github                             # GitHub configuration directory
│   └── workflows                       # GitHub Actions workflows
│       ├── integration-tests.yml       # Workflow for integration tests
│       ├── unit-tests.yml              # Workflow for unit tests
│       └── venom-tests.yml             # Workflow for venom tests
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
│   ├── Dockerfile                      # Dockerfile     - bFF
│   ├── requirements.txt                # Requirements   - python dependencies
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
├── sql                                 # SQL files      - database schema
│   ├── Dockerfile                      # Dockerfile     - sql                   
│   ├── schemadump.sql                  # Schema creation and sample data
│   └── schema.sql                      # Schema creation only
├── src                                 # Source code
│   └── __init__.py                     # Initialization and configuration
├── .env                                # Environment variables file                                
├── .gitignore                          # Git ignore     - version control
├── Makefile                            # Building and running
├── README.md                           # Project documentation
└── docker-compose.yml                  # Docker Compose configuration
```
<br>

### **Components:**
<br>

    Backend  (Go)           : Handles user data management, authentication, and PDF generation.
    Frontend (Python,Flask) : Provides the web interface for user interaction.
    Database (SQL)          : Stores user information.

<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>

### **Prerequisites:**
<br>

- `Go                        `: _Backend development_
- `Flask                     `: _Frontend development_
- `MySQL database            `: _Storing user data and templates_
- `wkhtmltopdf               `: _PDF generation_
- `Git                       `: _Version control_
- `Docker & Docker Compose   `: _Containerized deployment_


<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>

### **Environment variables file: .env file**
<br>

```sh
MYSQL_ROOT_PASSWORD=***
MYSQL_PASSWORD= ***
MYSQL_ROOT_USER=root
MYSQL_USER= ***
MYSQL_ADDR=cv_db-container
MYSQL_HOST=localhost
MYSQL_DATABASE=users
MYSQL_PORT=3306
```

<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>

### **Install Basic Tools:**
<br>

```sh
sudo apt update && sudo apt upgrade && sudo apt install -y git curl build-essential golang-go python3 python3-pip wkhtmltopdf docker.io docker-compose selinux-utils curl mysql-server
sudo mysql_secure_installation
pip install --break-system-packages Flask Flask-Bcrypt Flask-Migrate Flask-SQLAlchemy
```

<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>

### **Replace Paths**
<br>

```sh
PathCvProject="/bcn/github/CV_project"
grep -q "PathCvProject=" ~/.bashrc || echo "export PathCvProject=\"$PathCvProject\"                                         # Set path to CV project." >> ~/.bashrc && source ~/.bashrc
```

<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>

### **Create DB, users table**
<br>

```sh
sudo mysql -u root -p
CREATE DATABASE IF NOT EXISTS users;
USE users;
SOURCE /bcn/github/CV_project/sql/schemadump.sql;
```

<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>

### **Verify successful import**
<br>

```sh
mysql -u root -p users
SHOW DATABASES;
SHOW TABLES;
USE users;
DESCRIBE template;
DESCRIBE users;
SELECT * FROM template;
SELECT * FROM users;
```

<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>

### **Change DB user password**
<br>

```sh
ALTER USER 'CV_user'@'localhost' IDENTIFIED BY 'Y0ur_strong_password';
```

<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>

### **Build the backend API**
<br>

```sh
cd $PathCvProject/api
go mod tidy
go build -o bin/user-service main.go
export DB_USER="root"
export DB_PASSWORD="?????????????"
./bin/user-service
```

<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>

## BFF Flask app setup frontend
<br>

```sh
cd $PathCvProject/bff
python3 app.py -i 127.0.0.1 -p 8080
```

<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>


### **Github SSH conection**
<br>

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

<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>

### **Github commit**
<br>

```sh
echo "Enter commit message (Title case, infinitive verb, brief and clear summary of changes):"
read -p "CommitMssg: - " CommitMssg
cd "$PathCvProject" || exit
git add .
git commit -m "$CommitMssg"
git pull && git push origin main
```

<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>

### **Start the project: Makefile**
<br>

```sh
cd $PathCvProject && make
```

<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>


### ****Docker****


|   | Command | Description |
|:---|:---|:---|
|🔨 |`docker-compose build `  |# build
|🔼 |`docker-compose up    `  |# start
|🔘 |`docker-compose up -d `  |# run background
|🛑 |`docker-compose stop  `  |# only stop
|🔽 |`docker-compose down  `  |# stops and removes containers
|🩺 |`docker-compose ps    `  |# view running containers
|🧹 |`docker-compose rm    `  |# removes stopped service containers


<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>

### **Browser links**
<br>

- [Miro](https://miro.com/app/board/uXjVK6HA_1A=/)

- [template1](http://127.0.0.1:5000/template1)

- [template2](http://127.0.0.1:5000/template2)

- [template3](http://127.0.0.1:5000/template3)
#
![](https://i.imgur.com/waxVImv.png)

### **List of Free Learning Resources In Many Languages**
<br>

<div align="center" markdown="1">

[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)&#160;
[![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/)&#160;
[![Hacktoberfest 2023 stats](https://img.shields.io/github/hacktoberfest/2023/EbookFoundation/free-programming-books?label=Hacktoberfest+2023)](https://github.com/EbookFoundation/free-programming-books/pulls?q=is%3Apr+is%3Amerged+created%3A2023-10-01..2023-10-31)

</div>


### **Where To Look For Further Info :thinking:**
<br>

- [freeCodeCamp Guide](https://guide.freecodecamp.org/)
- [GeeksForGeeks](https://www.geeksforgeeks.org/)
- [Dev.To](https://dev.to/)
- [Stack Overflow](https://stackoverflow.com/)
- [Dzone](https://dzone.com/)  
<br>

### **Other free resources**
<br>

+ [English, Books By Programming Language](https://github.com/EbookFoundation/free-programming-books/blob/main/books/free-programming-books-langs.md/)
+ [English, Interactive Programming Resources](https://github.com/EbookFoundation/free-programming-books/blob/main/more/free-programming-interactive-tutorials-en.md)
+ [English Courses](https://github.com/EbookFoundation/free-programming-books/blob/main/courses/free-courses-en.md)

<div align="right">
  <b><a href="#table-of-contents">↥ Top 🔝</a></b>
</div>

### **Coding Practice Sites :zap:**
<br>

- :link: [CodeForces](http://codeforces.com/)
- :link: [CodeChef](https://www.codechef.com)
- :link: [Coderbyte](https://coderbyte.com/)
- :link: [CodinGame](https://www.codingame.com/)
- :link: [Cs Academy](https://csacademy.com/)
- :link: [HackerRank](https://hackerrank.com/)
- :link: [Spoj](https://spoj.com/)
- :link: [HackerEarth](https://hackerearth.com/)
- :link: [TopCoder](https://www.topcoder.com/)
- :link: [Codewars](https://codewars.com/)
- :link: [Exercism](http://www.exercism.io/)
- :link: [CodeSignal](https://codesignal.com/)
- :link: [Project Euler](https://projecteuler.net/)
- :link: [LeetCode](https://leetcode.com/)
- :link: [Firecode.io](https://www.firecode.io/)
- :link: [InterviewBit](https://www.interviewbit.com/)
- :link: [uCoder](https://ucoder.com.br)
- :link: [LintCode](https://www.lintcode.com/)
- :link: [CodeCombat](https://codecombat.com/)
- :link: [InterviewCake](https://www.interviewcake.com/)
- :link: [At Coder](https://atcoder.jp/)
- :link: [Codility](https://www.codility.com/)
- :link: [ICPC Problem Archive](https://icpc.kattis.com/)
- :link: [Codemia](https://codemia.io/)


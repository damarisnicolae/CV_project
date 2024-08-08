# Variables
include .env
# Targets
.PHONY: all db setup-db build-backend run-backend setup-frontend open-browser

all: db setup-db build-backend setup-frontend run-backend open-browser

db:
	@echo "Setting up MySQL database and user..."
	# mysql -u root -p"$(MySqlRootPass)" -e "CREATE DATABASE IF NOT EXISTS users;"
	# mysql -u root -p"$(MySqlRootPass)" -e "CREATE USER IF NOT EXISTS '$(DBUSER)'@'localhost' IDENTIFIED BY '$(DBPASS)';"
	# mysql -u root -p"$(MySqlRootPass)" -e "GRANT ALL PRIVILEGES ON users.* TO '$(DBUSER)'@'localhost'; FLUSH PRIVILEGES;"

setup-db:
	@echo "Updating DB schema..."
	mysql -u '$(DBUSER)' -p'$(DBPASS)' -e "USE cv_project; DROP TABLE IF EXISTS user, template;"
	mysql -u '$(DBUSER)' -p'$(DBPASS)' cv_project < $(PathCvProject)/sql/schemadump.sql

build-backend:
	@echo "Building backend API..."
	cd $(PathCvProject)/api && go build -o CV_project main.go

setup-frontend:
	@echo "Setting up frontend Flask app..."
	export FLASK_APP=app FLASK_ENV=development && cd $(PathCvProject)/bff && python3 app.py -i 127.0.0.1 -p 8080 &

run-backend:
	@echo "Running backend API..."
	export DBUSER="cv_user" DBPASS="Y0ur_strong_password" && cd $(PathCvProject)/api && ./CV_project &

open-browser:
	@echo "Opening browser with URL..."
	sleep 2  
	xdg-open http://127.0.0.1:5000/template1




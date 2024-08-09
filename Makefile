# Variables
include .env
# Targets
.PHONY: all db setup-db build-backend run-backend setup-frontend open-browser

all: db setup-db build-backend setup-frontend run-backend open-browser

db:
	@echo "Setting up MySQL database and user..."
	# mysql -u root -p"$(MYSQL_ROOT_PASSWORD)" -e "CREATE DATABASE IF NOT EXISTS users;"
	# mysql -u root -p"$(MYSQL_ROOT_PASSWORD)" -e "CREATE USER IF NOT EXISTS '$(MYSQL_USER)'@'localhost' IDENTIFIED BY '$(MYSQL_PASSWORD)';"
	# mysql -u root -p"$(MYSQL_ROOT_PASSWORD)" -e "GRANT ALL PRIVILEGES ON users.* TO '$(MYSQL_USER)'@'localhost'; FLUSH PRIVILEGES;"

setup-db:
	@echo "Updating DB schema..."
	mysql -u '$(MYSQL_USER)' -p'$(MYSQL_PASSWORD)' -e "USE cv_project; DROP TABLE IF EXISTS CV_user, template;"
	mysql -u '$(MYSQL_USER)' -p'$(MYSQL_PASSWORD)' cv_project < $(PathCvProject)/sql/schemadump.sql

build-backend:
	@echo "Building backend API..."
	cd $(PathCvProject)/api && go build -o CV_project main.go

setup-frontend:
	@echo "Setting up frontend Flask app..."
	cd $(PathCvProject)/bff && python3 app.py -i 127.0.0.1 -p 8080 &

run-backend:
	@echo "Running backend API..."
	export MYSQL_USER='$(MYSQL_USER)' MYSQL_PASSWORD='$(MYSQL_PASSWORD)' && cd $(PathCvProject)/api && ./CV_project &

open-browser:
	@echo "Opening browser with URL..."
	sleep 2  
	xdg-open http://127.0.0.1:5000/template1




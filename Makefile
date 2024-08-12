# Variables
include .env

# Targets
.PHONY: all path db api bff browser kill

all: path db api bff browser

export

path:
	@echo "\n * * * Setup...\n"
	PathCvProject=$(pwd)
	grep -q "PathCvProject=" ~/.bashrc || echo "export PathCvProject=\"$PathCvProject\"                                         # Set path to CV project." >> ~/.bashrc && source ~/.bashrc

db:
	@echo "\n * * * DB...\n "
	service mysql start && service mysql status && mysql -u root -p$(MYSQL_ROOT_PASSWORD) -e "CREATE DATABASE IF NOT EXISTS $(MYSQL_DATABASE);"
	mysql -u root -p$(MYSQL_ROOT_PASSWORD) -e "USE $(MYSQL_DATABASE); SELECT 1 FROM users LIMIT 1;" | grep -q 1 || mysql -u root -p$(MYSQL_ROOT_PASSWORD) $(MYSQL_DATABASE) < $(PathCvProject)/sql/schemadump.sql

api:
	@echo "\n * * * API...\n"
	@while ! nc -z 127.0.0.1 3306; do sleep 1; done
	cd $(PathCvProject)/api && MYSQL_USER=$(MYSQL_ROOT_USER) MYSQL_PASSWORD=$(MYSQL_PASSWORD) MYSQL_ADDR=127.0.0.1 go run main.go &

bff:
	@echo "\n * * * BFF...\n"
	@while ! nc -z 127.0.0.1 8080; do sleep 1; done
	cd $(PathCvProject)/bff && /usr/bin/python3 app.py -i 127.0.0.1 -p 8080 &

browser:
	@echo "\n * * * URL...\n"
	@while ! nc -z 127.0.0.1 5000; do sleep 1; done
	xdg-open http://127.0.0.1:5000

kill:
	@echo "\n * * * Kill...\n"
	@lsof -ti :8080 | xargs -r kill -9
	@lsof -ti :5000 | xargs -r kill -9
	@lsof -ti :3306 | xargs -r kill -9

dokcer:
	@echo "\n * * * Docker...\n"
	docker-compose up -d

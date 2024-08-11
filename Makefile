# Variables
include .env

# Targets
.PHONY: all db api bff browser kill

all: db api bff browser

db:
	@echo "\n * * * DB...\n "
	service mysql start

export

api:
	@echo "\n * * * API...\n"
	cd $(PathCvProject)/api && MYSQL_USER=$(MYSQL_USER) MYSQL_PASSWORD=$(MYSQL_PASSWORD) go run $(PathCvProject)/api/main.go &
	# MYSQL_USER=$MYSQL_USER MYSQL_PASSWORD=$MYSQL_PASSWORD go run main.go 

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


include .env
export

.PHONY: all DB API BFF

all: DB API BFF 

DB:
	@sudo systemctl start mysql

API:
	@until nc -z 127.0.0.1 3306; do sleep 1; done
	@cd $(PathCvProject)/api && MYSQL_USER=$(MYSQL_ROOT_USER) MYSQL_PASSWORD=$(MYSQL_ROOT_PASSWORD) MYSQL_ADDR=localhost MYSQL_PORT=3306 go run main.go &

BFF:
	@while ! nc -z 127.0.0.1 8080; do sleep 1; done
	@cd $(PathCvProject)/bff && /usr/bin/python3 app.py -i 127.0.0.1 -p 8080

# Docker
ON:
	@docker-compose -f ./docker-compose.yml up -d
OFF:
	@docker-compose down --rmi all -v --remove-orphans 
	@docker system prune -a -f || true

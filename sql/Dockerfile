FROM mysql:latest

WORKDIR /sql/

RUN apt-get update && apt-get install -y wget \
    && wget https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-v0.6.1.tar.gz \
    && rm dockerize-linux-amd64-v0.6.1.tar.gz \
    && apt-get remove -y wget \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/*

COPY . .


RUN docker run -d --name db_container \
    --restart always \
    --env-file ../.env \
    -p 3306:3306 \
    -v $(PWD)/schemadump.sql:/docker-entrypoint-initdb.d/dump.sql

CMD ["mysqld"]

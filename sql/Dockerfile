FROM mysql:8.0

COPY ./schemadump.sql /docker-entrypoint-initdb.d/dump.sql

COPY . .

CMD ["mysqld"]

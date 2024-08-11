FROM python:3.11-slim

ENV DB_HOST  cv_db-container
ENV DB_PORT  3306
ENV API_HOST cv_api-container
ENV API_PORT 8080

RUN apt-get update && apt-get install -y wget \
    && wget https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-v0.6.1.tar.gz \
    && rm dockerize-linux-amd64-v0.6.1.tar.gz \
    && apt-get remove -y wget \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .

RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 5000

CMD ["dockerize", "-wait", "tcp://${API_HOST}:${API_PORT}", "-timeout", "60s", "python", "app.py", "-i", "${API_HOST}", "-p", "${API_PORT}"]

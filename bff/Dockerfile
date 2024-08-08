FROM python:3.11-slim

WORKDIR /bff

COPY bff/requirements.txt /bff/
RUN pip3 install --upgrade pip && \
    pip3 install -r requirements.txt

COPY bff/ /app/

EXPOSE 5000

CMD ["python3", "app.py", "-i", "127.0.0.1", "-p", "5000"]


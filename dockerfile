FROM python:3.9-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --upgrade pip
RUN pip install Flask prometheus_client
RUN pip install --no-cache-dir -r requirements.txt
COPY aplication/app.py .
CMD [ "python","app.py" ]
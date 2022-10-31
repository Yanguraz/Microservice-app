FROM python:3.10

LABEL islamianguraz "islamyang123@gmail.com"

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

WORKDIR /app
COPY . /app

RUN pip install --upgrade pip
RUN pip install -r /app/requirements.txt

CMD ["flask", "run", "--host=0.0.0.0"]
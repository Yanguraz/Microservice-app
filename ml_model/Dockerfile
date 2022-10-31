FROM python:3.10

LABEL islamianguraz "islamyang123@gmail.com"

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

WORKDIR /ml
COPY . /ml

RUN pip install --upgrade pip
RUN pip install -r /ml/requirements.txt

CMD ["python", "./main.py", "--host=0.0.0.0"]
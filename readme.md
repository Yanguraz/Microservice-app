# Microsevice-app

## A simple NLP intent recognition microservice web-application built using gRPC, Python, and Go.

**How to run this repository locally**
- Run all of the services via docker-compose file.
  - `main service`: run `python app.py`
  - `auth service`: run `go run main.go`
  - `predict service`: run `go run main.go`
  - `ml service`: run `go run main.py`

**Docker commands** 
  - main-run:
	`docker-compose up`

  - main-start:
	`docker-compose up --build`

  - main-stop:
	`docker-compose down`

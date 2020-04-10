## About

This repo contains 2 simple backend apps that provide several useful API:
- Auth backend app [Go Simple App](go-simple-app)
    - Register new user
    - Login
    - Get active user info
- Resources backend app [Py Simple App](py-simple-app)
    - Fetch resources
    - Aggregate resources price
    - Get active user info

If you want to know about the detail of each API and want to try it, you can refer to swagger docs:
- [Auth backend app](go-simple-app/swagger.yaml)
- [Resources backend app](py-simple-app/swagger.yaml)

## Authentication
This apps use Json Web Token (JWT) as an authentication method for the API. 
Authentication method is done by Auth backend app. Any authentication in Resources backend app is also done by Auth backend app, so you have to run Resources backend app along with Auth backend app.

## How to run
#### First method
This apps use docker compose to setup the environment. So, you can run this apps by executing ```docker-compose up```.
#### Second method
You can also run each apps individually using docker:
- Auth backend app: 
```
$ docker build -t go-simple-app .
$ docker run go-simple-app
```
- Resources backend app:
```
$ docker build -t py-simple-app .
$ docker run py-simple-app
```
#### Third method
You can also run each apps individually from sources:
- Auth backend app: 
    - Install go
```
$ cd go-simple-app
$ go mod tidy
$ go run main.go
```
- Resources backend app:
    - Install python
```
$ cd py-simple-app
$ pip install -r requirements.txt
$ python -u src/main.py
```

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

## How to run
This apps use docker compose to setup the environment. So, you can run this apps by executing ```docker-compose up```.

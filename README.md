## AUTH SERVICE

### Description
This repo contains project that handle auth-related services. 
This service is part of a big system. 
The whole system will be used to present technology show case.

### Features
- Register user
- Login
- Get active user

This service implement JWT as auth method.
This service serve feature that is mentioned above through HTTP.

### How to run
#### Docker
- Install docker
- Build and run docker image as below
```shell script
$ docker build -t auth-service .
$ docker run -p 8080:8080 auth-service
```

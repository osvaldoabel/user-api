# Documentation

[![Build Status](./build-passing.svg)]()

. [Introduction](#introduction)  
. [Project Requirements](#requirements)  
. [Api Endoints](#api-endpoints)  
. [TO DO](#to-do)  


# Introduction
 In this repository, we have everything we need to run this **User Api** project using [docker containers](http://docker.com). Below you can see a basic guide to learn how to run it in your local environment.

In this project we create a *REST API* using [hexagonal archiecture or ports & adpters architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)) resulting in a robust and flexible project. So it's  pretty simple to add new features, new Entities, etc. Including if we need to add a **console adapter**, the **application layers** will keep intact.

### Requirements

You just need to have [Git ](https://docs.docker.com/install) and [Docker ](https://docs.docker.com/install) installed and running on your machine.

#### Clone Project

```bash
$ git clone https://gitlab.com/osvaldoabel/user-api project-name
```
#### Run Project

```bash
# Get into the cloned repo
$ cd user-api

# Build and start containers, 
$ docker-compose up -d --build 
```

```bash
# You can also verify the containers status
$ docker-compose ps
```

```bash
  # get into vma.app to do anything you might need
$ docker exec -it vm.app bash
```


## API Endpoints 
- #### CREATE User

```
POST http://localhost:8800/users HTTP/1.1
Content-Type: application/json

{
    "name": "User example",
    "age": 40,
    "email": "abel44@test.com",
    "password": "123456",
    "status": "active",
    "address":  "Rua teste, 123"
}
```

- #### UPDATE User

```
PUT http://localhost:8800/users/4d4edebc-ed56-4fa7-bb14-df0f9204a16a HTTP/1.1
Content-Type: application/json
Authorization: Bearer {{token}}

{
    "id": "4d4edebc-ed56-4fa7-bb14-df0f9204a16a",
    "name": "user ",
    "age": 37,
    "address": "Rua teste, 123 ate 9"
}
```
- #### LIST users (Paginated)

```
# /v1/users?per_page=10

GET http://0.0.0.0:8800/users?page=4&limit=2 HTTP/1.1
Content-Type: application/json 
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTIwNTM1NDgsInN1YiI6ImJjNTkwMzg1LTFjOTgtNDNkNy1iN2FmLTY3OWI1OGUzNzc5YSJ9.oJhTOIGNqYZ0Wb4rwrviVlKOjoesvqg0OtETn3y0Nsg

```

- #### SHOW User

```
#/v1/users/fa9f88f4-4fe8-46d8-afb3-85886c50ec4c 

curl -X GET http://localhost:8888/v1/users/fa9f88f4-4fe8-46d8-afb3-85886c50ec4c -H 'Content-Type: application/json' \
  -d '{
    "name": "User 3 - updated",
    "email": "updated3@example.com",
    "status": "active",
    "address": "My Address",
    "password": "semsenha",
    "age": 35
}'
```

- #### DELETE User

```
#/v1/users/fa9f88f4-4fe8-46d8-afb3-85886c50ec4c 

curl -X DELETE http://localhost:8888/v1/users/fa9f88f4-4fe8-46d8-afb3-85886c50ec4c -H 'Content-Type: application/json'
```

# To DO
- Caching (With [redis](https://redis.io/) )
- Authentication / Authorization
- More [Advanced Logging system]()

**NOTE:** It Would be interesting if we persisted our logs into Elasticsearch. 
 **pros**:  
 . You can serve as many microservice as you'll need
 . You'll have an Asynchronous system (very good most of the time. )
 . etc.  
 **Cons**: 
 . Increases project complexity

 **NOTE 2**
This is just a hypothetical scenario. 
If your project will not grow to this magnitude, you don't need to implement it this way.

Developed by [Osvaldo Abel](https://gitlab.com/osvaldoabel)


# Go Project Template

A simple go project structure setup with all dependencies you need to get your MVP off the ground.

The project is suited as starting point for simple webapp MVP. Running the project starts the app wrapped using the go-shipped
webserver.

**Features**

* Live reloading
* sqllite3 database for local development
* ORM ready to go
* A simple database migration system
* Injection of environment variables into config files
* Simple spring-like repository structure for convenient database queries

## Installation

```bash
git clone https://github.com/eldonaldo/go-project-template
cd go-project-template
go get -u github.com/cosmtrek/air
```

## Run Project

```bash
cd go-project-template
air
```

Then interact with the app using the command line 

```bash
$ curl localhost:8008/create?name=John
$ User with name John1 created

$ curl localhost:8008/greet?name=John
$ Hi John
$ {
$   "ID": 1,
$   "CreatedAt": "2021-04-07T21:32:14.135761+02:00",
$   "UpdatedAt": "2021-04-07T21:32:14.135761+02:00",
$   "DeletedAt": null,
$   "Name": "John"
$ }
```

# Project Setup

* _cmd_: Contains the apps binaries
* _core_: The app's core functionality resides here (such as database handling, migrations etc.)
* _scripts_: Automation scripts. There is a script to bootstrap a new migration and another to downgrade an already
applied migration.
* _server_: Handles the HTTP server setup
* _user_: A user module

## Libraries Used

* [github.com/caarlos0/env/v6 - Read env variables into struct fields](https://github.com/caarlos0/env/v6)
* [github.com/pressly/goose - Database migrations](https://github.com/pressly/goose)
* [gorm.io/gorm - ORM](https://gorm.io/gorm)
* [github.com/cosmtrek/air - Live reloading](https://github.com/cosmtrek/air)
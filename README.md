# A Simple Go Project Template - Suited for Webapp MVPs

A simple go project structure setup with all dependencies you need to get your MVP off the ground :rocket:

The project is suited as starting point for a simple webapp MVP. Besides the core files needed to get the setup running
it contains a simple user module showcasing the intended project structure. Running the project starts the app exposing
a REST API using the go-shipped webserver.

**Features**

* Live reloading :trollface:
* sqllite3 database for local development :heart:
* ORM ready to go :runner:
* A simple database migration system :raised_hands:
* Injection of environment variables into config files :notes:
* Simple spring-like repository structure for convenient database queries :star:


I hope you find it useful! :beer: :pizza:

## Installation

```bash
git clone https://github.com/eldonaldo/go-project-template
cd go-project-template

# You need to get air separately to get the live reloading running 
go get -u github.com/cosmtrek/air
```

## Run Project

```bash
cd go-project-template

# Runs the project with live reloading enabled
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

## Project Setup

* _cmd_: Contains the app binaries
* _core_: The app's core functionality resides here (such as database handling, migrations etc.)
* _core/db/migrations_: SQL migration files go in here. Use `./scripts/migration_create.sh migration_name` to create a
  new one.
* _scripts_: Automation scripts. There is a script to bootstrap a new migration and another to downgrade an already
  applied migration.
* _server_: Exposes the app as REST API and handles the HTTP server setup.
* _user_: A user module showcasing a simple structure using repository and REST handler.

## Libraries Used

* [github.com/caarlos0/env - Read env variables into struct fields](https://github.com/caarlos0/env)
* [github.com/pressly/goose - Database migrations](https://github.com/pressly/goose)
* [gorm.io/gorm - ORM](https://gorm.io/)
* [github.com/cosmtrek/air - Live reloading](https://github.com/cosmtrek/air)

:wave:
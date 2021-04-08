# A Simple Go Project Template - Suited for Webapp MVPs

A simple go project structure setup with all dependencies you need to get your MVP off the ground. :rocket:

The project is suited as starting point for a simple webapp MVP. Besides the core files needed to get the setup running
it contains a simple user module showcasing the intended project structure. Running the project starts the app exposing
a REST API using the go-shipped webserver.

**Features**

* Hot reloading. :trollface:
* sqllite3 database for local development. :heart:
* ORM ready to go. :runner:
* A simple database migration system. :raised_hands:
* Injection of environment variables into config files. :notes:
* Simple spring-like repository structure for convenient database queries. :star:
* Testing setup. :construction_worker:
  
Feedback and PRs welcome! I hope you find it useful. :beer: :pizza:

## Installation and Run Project

```bash
$ git clone https://github.com/eldonaldo/go-project-template
$ cd go-project-template

# You need to get air separately to get the hot reloading running 
$ go get -u github.com/cosmtrek/air

# Runs the project with hot reloading enabled
$ air
```

Interact with the app using the command line :computer:

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
  applied migration. Migrations in this folder are automatically executed once in order upon app start.
* _server_: Exposes the app as REST API and handles the HTTP server setup.
* _user_: A user module showcasing a simple structure using repository and REST handler.

### Changing the Project Name

To change the project name from _go-project-name_ to something else you need to edit the _go.mod_ module end edit the
first line `module github.com/eldonaldo/go-project-template` to `module github.com/your-username/your-new-name`. Imports
in all files need to be changed according (your IDE probably can do that for you :recycle:). Further, you might also
want change `cmd/project-name` and therefore you also need to change line 6 of `.air.conf`. But that should be it
then. :white_check_mark:

## Libraries Used

* [github.com/caarlos0/env - Read env variables into struct fields](https://github.com/caarlos0/env)
* [github.com/pressly/goose - Database migrations](https://github.com/pressly/goose)
* [gorm.io/gorm - ORM](https://gorm.io/)
* [github.com/cosmtrek/air - Hot reloading](https://github.com/cosmtrek/air)
* [github.com/stretchr/testify/assert - For convenient testing](github.com/stretchr/testify/assert)

:wave:
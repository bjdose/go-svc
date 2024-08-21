## Running the project
From the root level of the project folder, execute this command (this assumes that you have
[GNU make](https://www.gnu.org/software/make/) and a recent version
of [Docker](https://www.docker.com/products/docker-desktop) installed on your machine):

~~~
make up_build 
~~~

If the code has not changed, subsequent runs can just be `make up`.

Then start the front end:

~~~
make start
~~~

Hit the front end with your web browser at `http://localhost:80`. You can also access a web
front end to the logger service by going to `http://localhost:8082` (or whatever port you
specify in the `docker-compose.yml file`).

To stop everything:

~~~
make stop
make down
~~~

While working on code, you can rebuild just the service you are working on by
executing

`make build_auth`

Where `build_auth` is one of the services:

- build_broker
- build_auth
- build_logger
- build_mail
- build_listener
- build_frontend

All make commands:

~~~
Choose a command:

up               starts all containers in the background without forcing build
down             stop docker compose
up_build         stops docker-compose (if running), builds all projects and starts docker compose
build_auth       builds the authentication binary as a linux executable
build_logger     builds the logger binary as a linux executable
build_broker     builds the broker binary as a linux executable
build_listener   builds the listener binary as a linux executable
build_mail       builds the mail binary as a linux executable
start            starts the front end
stop             stop the front end
~~~
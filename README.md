# 개인프로젝트 (Simple SNS)

- user-service (current project)

## Running App

1. copy .env file

~~~bash
$ cp dist.env .env
~~~

2. fill the env variables

~~~bash
SERVER_PORT=8080
...
~~~

3. install dependencies

~~~bash
$ go get
~~~

4. run

~~~bash
$ make run
~~~

## Running App using Docker-Compose

1. copy .env file

~~~bash
$ cp dist.env .env
~~~

2. fill the env variables

~~~bash
SERVER_PORT=8080
...
~~~

3. run

~~~bash
$ docker-compose up
~~~
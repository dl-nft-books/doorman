# doorman

## Description

Service for generating access and refresh JWTs

## Install

  ```bash
  git clone doorman
  cd 
  go build main.go
  export KV_VIPER_FILE=./config.yaml
  ./main migrate up
  ./main run service
  ```

## Documentation

We do use openapi:json standard for API. We use swagger for documenting our API.

To open online documentation, go to [swagger editor](http://localhost:8080/swagger-editor/) here is how you can start it
```bash
  cd docs
  npm install
  npm start
```
To build documentation use `npm run build` command,
that will create open-api documentation in `web_deploy` folder.

To generate resources for Go models run `./generate.sh` script in root folder.
use `./generate.sh --help` to see all available options.


## Running from docker 
  
Make sure that the docker is installed.
Use `docker run ` with `-p 8080:80` to expose port 80 to 8080

```bash
docker build -t doorman .
docker run -e KV_VIPER_FILE=/config.yaml doorman
```

## Running from Source

* Set up environment value with config file path `KV_VIPER_FILE=./config.yaml`
* Provide valid config file
* Launch the service with `migrate up` command to create database schema
* Launch the service with `run service` command


### Database
For services, we do use ***PostgresSQL*** database. 
You can [install it locally](https://www.postgresql.org/download/) or use [docker image](https://hub.docker.com/_/postgres/).

## Contacts

Dmytro Zakharov is responsible for this service. Can be contacted via the _Telegram_ (`@ZamDimon`)

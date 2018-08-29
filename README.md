# PNGR
dockerized (postgres + nginx + golang + react) starter kit

I've only implemented basic user management stuff in terms of actual features, but this scaffolding can be extended to serve a huge variety of purposes.

This project is meant to be a starting point. Feel free to create issues with suggestions, or pull requests for security improvements or developer ergonomics.

## Requirements
- Install docker && docker-compose

## Start the Dev Server
1) `sudo docker-compose up`
2) Visit https://localhost (and approve the self-signed cert)
3) Make changes to either golang or react code, and watch the app rebuild itself!

## Production Builds
*Warning: this code is pre-alpha quality! Run in production at your own risk*

- Generate a production container with `sudo docker build .` 
- You will need to setup nginx in production for SSL termination and port forwarding to `:5000` look at `nginx/nginx.prod.conf` for ideas on how to do this.

--- 

## Postgres
Some tips for working with your postgres docker instance

### Creating and running migrations
Migrations are run using [pgmigrate](https://github.com/yandex/pgmigrate).

- `postgres/new-migration.sh my_migration_name` will create a template for the next migration-
- `postgres/run-migrations.sh` will execute any new migrations 

### Opening a psql client
`sudo docker exec -it pngr_postgres_1 psql --username postgres --dbname postgres`
Remember to use `\q` to exit.

### Rebuilding your database from scratch
If you want to clear out your postgres instance and start fresh, due to a bad migration or some other issue, normal container recreation isn't enough, because docker compose creates a volume for postgres data.

To *fully* reset your postgres instance, run:
`sudo docker-compose down -v && sudo docker-compose up --build --force-recreate`

--- 

## Nginx
Nginx is simply used to route requests to the front-end and back-end based on path.
It also terminates SSL so that we don't have to deal with certs in our app layer.

--- 

## Golang
Almost-vanilla golang api:
- [dep](https://github.com/golang/dep) for dependencies (to be replaced with go modules with Go 1.11)
- [Sqlx](https://github.com/jmoiron/sqlx) for cleaner interactions with postgres

--- 

## React
The basic building blocks of the front-end are:
- [Create React App](https://github.com/facebookincubator/create-react-app)
- [React Router](https://github.com/ReactTraining/react-router)
- [Unstated](https://github.com/jamiebuilds/unstated)
- [Semantic UI React](https://react.semantic-ui.com/)

# PNGR Stack
Dockerized (postgres + nginx + golang + react) starter kit

This repo only implements basic user signup, session management, and a toy `post` type to demonstrate basic CRUD.
The idea is that this scaffolding can be forked and extended to serve a _huge_ variety of purposes.

PNGR is _not_ a CMS.

Feel free to create issues with suggestions, or pull requests for security or developer ergonomics improvements.

## Requirements
- Install docker && docker-compose

## Quick Start
1) `sudo docker-compose up`
2) Visit https://localhost (and approve the self-signed cert)
3) Make changes to either golang or react code, and enjoy hot-reload goodness!

## Deploying to Production
*Warning: this code is pre-alpha quality! Run in production at your own risk*

This part is only sketched out, as everyone's production deployment will look different:
- **P** Consider running an actual postgres instance. Running a production database in docker makes me sweat.
- **N** Look at `nginx/nginx.prod.conf` for ideas on what a production configuration might look like.
- **G** Use `golang/Dockerfile.prod`
- **R** Use `react/Dockerfile.prod`
	- e.g. From project root you can run `sudo docker build -t react-prod -f react/Dockerfile.prod react` 
	- Test it out with `sudo docker run --net=host react-prod` then hit `http://localhost` in your browser

--- 

## Postgres
Some tips for working with your postgres docker instance

### Creating and running migrations
Migrations are run using [go-migrate](https://github.com/golang-migrate/migrate).

I put together little bash scripts to help you get stuff done.
- `sudo postgres/new-migration.sh my_migration_name` will create a template for the next migration.
- `sudo postgres/run-migrations.sh` will execute any new migrations 

You can do more advanced migrate commands 

### Opening a psql client
`sudo docker-compose exec postgres psql -U postgres`
Remember to use `\q` to exit.

### Rebuilding your database from scratch
If you want to clear out your postgres instance and start fresh, due to a bad migration or some other issue, normal container recreation isn't enough, because docker compose creates a volume for postgres data.

To *fully* reset your development environment:
`sudo docker-compose down -v && sudo docker-compose up --build --force-recreate`

--- 

## Nginx
Nginx is simply used to route requests to the front-end and back-end based on path.
It also terminates SSL so that we don't have to deal with certs in our app layer.

--- 

## Golang
Almost-vanilla golang api:
- Makes use of go modules for dependencies(!)
- [jwt-gp](github.com/dgrijalva/jwt-go) for JSON Web Tokens
- [sqlx](https://github.com/jmoiron/sqlx) for better postgres interface

--- 

## React
The basic building blocks of the front-end are:
- [Create React App](https://github.com/facebookincubator/create-react-app) (unejected!)
- [React Router](https://github.com/ReactTraining/react-router)
- [Unstated](https://github.com/jamiebuilds/unstated) for state management
- [Semantic UI React](https://react.semantic-ui.com/) as component library

# PNGR Stack
Dockerized (postgres + nginx + golang + react) starter kit

Only implements basic user signup, session management, and a toy `post` type to demonstrate basic CRUD. PNGR is _not_ a CMS.

## Features
- Hot-reload, front and back, including a test-runner for golang changes
- JSON Web-Token cookies with automatic refresh: ready for horizontal scaling
- Multi-stage builds for small production images
- A worker container for async (non-API) tasks
- Feature development is up to you!

## Requirements
Install `docker` && `docker-compose`

## Quick Start
1) `docker-compose up`
2) Visit `https://localhost` (*note **https***)
3) Approve the self-signed cert
4) Make changes to either golang or react code, and enjoy hot-reload goodness!

Preview of the app:
![Screenshot of the app](docs/demo.png?raw=true "Screenshot")

## Deploying to Production
*Warning: Run in production at your own risk!*

`docker-compose.prod.yml` is designed for a setup where postgresql is _not_ dockerized, as dockerizing the persistence layer makes it too easy to destroy your real data. Pulling images from a registry and/or using CI/CD is up to you.

```bash
# build lean images and run them in a detached state
docker-compose -f docker-compose.prod.yml up --build -d
```

Don't forget to copy `.env.example` -> `.env` and setup your secrets/passwords for the new environment!

--- 

## Postgres

### Creating and running migrations
Migrations are created and run using [go-migrate](https://github.com/golang-migrate/migrate).

```bash
# create a template for the next migration
postgres/new my_migration_name

# execute any new migrations (this is also run automatically the container is created)
postgres/migrate up

# go down 1 migration
postgres/migrate down 1
```

### Opening a psql client
```bash
docker-compose exec postgres psql -U postgres
```
Remember to use `\q` to exit.

### Rebuilding your dev environment, including database, from scratch
Maybe your postgres went sideways from a wonky migration and it's easier to restart from scratch.
```bash
docker-compose down -v && docker-compose up --build --force-recreate
```

## Nginx
Nginx is used to terminate SSL and route requests to the front-end and back-end based on path.

## Golang
- [jwt-go](https://github.com/dgrijalva/jwt-go) for JSON Web Tokens
- [sqlx](https://github.com/jmoiron/sqlx) for better postgres interface

## React
- [Create React App](https://github.com/facebookincubator/create-react-app) (unejected!)
- [React Context](https://reactjs.org/docs/context.html) for global state
- [React Router](https://github.com/ReactTraining/react-router)
- [Semantic UI React](https://react.semantic-ui.com/) for component library

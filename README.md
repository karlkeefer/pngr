## New in September 2021
I've finally upstreamed various upgrades I made on forked projects:

1) Auto-generated database code and mocks with the amazing [sqlc](https://github.com/kyleconroy/sqlc) and [gomock](https://github.com/golang/mock) packages 🙇‍♂️
2) Vastly simplified routing on [front-end](./react/src/Routes/Routes.js) and [back-end](./golang/server/routes.go)... thanks [httprouter](github.com/julienschmidt/httprouter)!
3) Functional-style components throughout, including some helpful [custom hooks to simplify building forms](./react/src/Routes/Posts/PostForm.js)
4) Added password reset functionality (and stubbed out a transactional email interface)
5) Standardized handler type for the API, with improved default middlewares
6) Simplified migrations and postgres helper scripts
7) Upgraded the Semantic UI integration to allow altering [theme variables](https://react.semantic-ui.com/theming/)
8) Various other improvements/tweaks/bugfixes
9) Added a ping-pong emoji to calcify my preferred pronunciation 😅

# PNGR Stack 🏓
Dockerized (postgres + nginx + golang + react) starter kit

Only implements `users`, `sessions`, and a toy `post` type to demonstrate basic CRUD. PNGR is _not_ a CMS.

## Features
- Hot-reload, front and back, including a test-runner for golang changes
- [golang-migrate](https://github.com/golang-migrate/migrate) already configured for easy migrations
- [sqlc](https://github.com/kyleconroy/sqlc) for auto-generated sql bindings and mocks (also rigged with hot-reload!)
- [jwt-go](https://github.com/dgrijalva/jwt-go) cookies with automatic refresh: ready for horizontal scaling
- A golang worker container stubbed out for async (non-API) tasks
- Unejected [Create React App](https://github.com/facebookincubator/create-react-app) as the basis for the front-end
- [React Router](https://github.com/ReactTraining/react-router) for simple front-end routing
- [React Context](https://reactjs.org/docs/context.html) for global state
- [Semantic UI React](https://react.semantic-ui.com/) for component library
- Feature development is up to you!

## Requirements
Install `docker` && `docker-compose`

## Quick Start
```bash
# clone the repo
git clone https://github.com/karlkeefer/pngr.git my_project_name

# copy the .env template for your local version
cp .env.example .env

# build and start the containers
docker-compose up
```
1) Visit `https://localhost` (*note **https***)
2) Approve the self-signed cert
3) Make changes to go, sql, or react code, and enjoy hot-reload goodness!

<img src="./docs/demo.png" width="400"/>

## Database Helpers

### Migrations
Migrations are created and run using simple wrappers around [go-migrate](https://github.com/golang-migrate/migrate).

```bash
# create files for a new migration
postgres/new my_migration_name

# execute any new migrations (this is also run automatically the container is created)
postgres/migrate up

# go down 1 migration
postgres/migrate down 1
```

### Open a psql client
```bash
# remember to use \q to exit
postgres/psql
```

## Rebuild everything, including database(!), from scratch
Maybe your postgres went sideways from a wonky migration and it's easier to restart from scratch.
```bash
docker-compose down -v && docker-compose up --build --force-recreate
```

## Deploy to Production
*Warning: Run in production at your own risk!*

`docker-compose.prod.yml` is designed for a setup where postgresql is _not_ dockerized. Pulling images from a registry and/or using CI/CD is up to you.

Don't forget to copy `.env.example` -> `.env` and setup your secrets/passwords for the new environment!
At minimum you'll need to change `ENV`, `APP_ROOT`, and `TOKEN_SECRET`!

```bash
# build production images, and run them in a detached state
docker-compose -f docker-compose.prod.yml up --build -d
```

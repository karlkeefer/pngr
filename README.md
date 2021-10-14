# PNGR Stack 🏓
Dockerized (postgres + nginx + golang + react) starter kit

Only implements `users`, `sessions`, `password_resets`, and a toy `post` type to demonstrate basic CRUD. PNGR is _not_ a CMS.

## Features
- Hot-reload, front and back, including a test-runner for golang changes
- [golang-migrate](https://github.com/golang-migrate/migrate) already configured for easy migrations
- [sqlc](https://github.com/kyleconroy/sqlc) for auto-generated sql bindings and [gomock](https://github.com/golang/mock) for auto-generated mocks (also rigged with hot-reload!)
- [jwt-go](https://github.com/dgrijalva/jwt-go) cookies with automatic refresh: ready for horizontal scaling
- Simple [default middleware for CORS, CSRF, cookie parsing, etc](./golang/server/middleware.go).
- A golang worker container stubbed out for async (non-API) tasks
- "Server-side rendering" with a [prerender sidecar container](./prerender/Dockerfile)
- Unejected [Create React App](https://github.com/facebookincubator/create-react-app) as the basis for the front-end
- Functional-style components throughout, including some helpful [custom hooks to simplify building forms](./react/src/Routes/Posts/PostForm.js)
- [React Router](https://github.com/ReactTraining/react-router) for [front-end routing](./react/src/Routes/Routes.js)
- [httprouter](github.com/julienschmidt/httprouter) for [simple back-end routing](./golang/server/routes.go)
- [React Context](https://reactjs.org/docs/context.html) for global state
- [Semantic UI React](https://react.semantic-ui.com/) for component library - allows changing [theme variables](https://github.com/Semantic-Org/Semantic-UI/blob/master/src/themes/default/globals/site.variables) with hot-reload
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

# goto a migration by index
postgres/migrate goto 3
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

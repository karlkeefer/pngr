# PNGR Stack 🏓
[![Build](https://github.com/karlkeefer/pngr/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/karlkeefer/pngr/actions/workflows/build.yml)

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
- [React Router](https://github.com/ReactTraining/react-router) for [front-end routing](./react/src/Routes/Routes.js)
- [httprouter](github.com/julienschmidt/httprouter) for [simple back-end routing](./golang/server/routes.go)
- Uses [React Context](https://reactjs.org/docs/context.html) for global user state
- Functional-style components throughout, including some helpful [custom hooks to simplify building forms](./react/src/Routes/Posts/PostForm.js)
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
docker-compose up # или используй ту-же комманду с флагом --build, чтобы изменения в коде вызывали ребилды (нужно при девелопе короче)


# To start the stack, defined by the Compose file in detached mode, run:
 docker-compose up --build -d
# Then, you can use docker-compose stop to stop the containers and docker-compose down to remove them.

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

## Run in Production
*Warning: Run in production at your own risk!*

`docker-compose.prod.yml` is designed for a setup where postgresql is _not_ dockerized.

Don't forget to copy `.env.example` -> `.env` and setup your secrets/passwords for the new environment!
At minimum, you'll need to change `ENV`, `APP_ROOT`, and `TOKEN_SECRET`!

```bash
# build production images, and run them in a detached state
docker-compose -f docker-compose.prod-build.yml up --build -d
```

Note: using your production server as your build server is a bad idea, so you should consider using a registry...

### Using CI
You can modify the github action to push built containers to a container registry. The containers are tagged with the commit SHA by default.

You will also need to update `docker-compose.prod.yml` to point to your container registry.

```bash
# pull containers from a registry using a tag, then run them in a detached state
SHA=2c25e862e0f36e0fc17c1703e4f319f0d9d04520 docker-compose -f docker-compose.prod.yml up -d
```
### Project Tree

```
pngr
├─ .git
│  ├─ config
│  ├─ description
│  ├─ FETCH_HEAD
│  ├─ fsmonitor--daemon
│  │  └─ cookies
│  ├─ HEAD
│  ├─ hooks
│  │  ├─ applypatch-msg.sample
│  │  ├─ commit-msg.sample
│  │  ├─ fsmonitor-watchman.sample
│  │  ├─ post-update.sample
│  │  ├─ pre-applypatch.sample
│  │  ├─ pre-commit.sample
│  │  ├─ pre-merge-commit.sample
│  │  ├─ pre-push.sample
│  │  ├─ pre-rebase.sample
│  │  ├─ pre-receive.sample
│  │  ├─ prepare-commit-msg.sample
│  │  ├─ push-to-checkout.sample
│  │  └─ update.sample
│  ├─ index
│  ├─ info
│  │  └─ exclude
│  ├─ logs
│  │  ├─ HEAD
│  │  └─ refs
│  │     ├─ heads
│  │     │  └─ master
│  │     └─ remotes
│  │        └─ origin
│  │           └─ HEAD
│  ├─ objects
│  │  ├─ info
│  │  └─ pack
│  │     ├─ pack-6cd532cb0f8623bd45b152a89cd6251d070656f0.idx
│  │     └─ pack-6cd532cb0f8623bd45b152a89cd6251d070656f0.pack
│  ├─ packed-refs
│  └─ refs
│     ├─ heads
│     │  └─ master
│     ├─ remotes
│     │  └─ origin
│     │     └─ HEAD
│     └─ tags
├─ .github
│  └─ workflows
│     └─ build.yml
├─ .gitignore
├─ docker-compose.ci.yml
├─ docker-compose.prod-build.yml
├─ docker-compose.prod.yml
├─ docker-compose.yml
├─ docs
│  └─ demo.png
├─ golang
│  ├─ .gitignore
│  ├─ cmd
│  │  ├─ server
│  │  │  └─ server.go
│  │  └─ worker
│  │     ├─ worker.go
│  │     └─ worker_test.go
│  ├─ db
│  │  ├─ db.go
│  │  ├─ mock.go
│  │  ├─ models.go
│  │  ├─ post.sql.go
│  │  ├─ querier.go
│  │  ├─ reset.sql.go
│  │  ├─ user.marshal.go
│  │  ├─ user.sql.go
│  │  └─ wrapper
│  │     ├─ mock.go
│  │     └─ wrapper.go
│  ├─ env
│  │  ├─ db.go
│  │  └─ env.go
│  ├─ errors
│  │  └─ errors.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ mail
│  │  ├─ html.go
│  │  ├─ html_test.go
│  │  ├─ logger.go
│  │  └─ mail.go
│  ├─ server
│  │  ├─ handlers
│  │  │  ├─ helpers.go
│  │  │  ├─ posts.go
│  │  │  ├─ resets.go
│  │  │  ├─ session.go
│  │  │  └─ user.go
│  │  ├─ helpers.go
│  │  ├─ jwt
│  │  │  ├─ jwt.go
│  │  │  └─ jwt_test.go
│  │  ├─ middleware.go
│  │  ├─ routes.go
│  │  ├─ server.go
│  │  └─ write
│  │     └─ write.go
│  ├─ server.Dockerfile
│  ├─ server.modd.conf
│  ├─ sql
│  │  ├─ queries
│  │  │  ├─ post.sql
│  │  │  ├─ reset.sql
│  │  │  └─ user.sql
│  │  └─ schema
│  │     ├─ 001_users.down.sql
│  │     ├─ 001_users.up.sql
│  │     ├─ 002_reset_password.down.sql
│  │     ├─ 002_reset_password.up.sql
│  │     ├─ 003_posts.down.sql
│  │     └─ 003_posts.up.sql
│  ├─ sqlc.yaml
│  ├─ worker.Dockerfile
│  └─ worker.modd.conf
├─ LICENSE
├─ nginx
│  ├─ cert.pem
│  ├─ Dockerfile
│  ├─ key.pem
│  ├─ nginx.conf
│  └─ nginx.prod.conf
├─ postgres
│  ├─ Dockerfile
│  ├─ init.sh
│  ├─ migrate
│  ├─ new
│  └─ psql
├─ prerender
│  ├─ Dockerfile
│  ├─ package.json
│  └─ server.js
├─ react
│  ├─ .dockerignore
│  ├─ .eslintrc
│  ├─ .gitignore
│  ├─ craco.config.js
│  ├─ Dockerfile
│  ├─ jsconfig.json
│  ├─ nginx.conf
│  ├─ package-lock.json
│  ├─ package.json
│  ├─ public
│  │  ├─ favicon.ico
│  │  ├─ index.html
│  │  └─ manifest.json
│  ├─ src
│  │  ├─ Api.js
│  │  ├─ App.js
│  │  ├─ index.css
│  │  ├─ index.js
│  │  ├─ Nav
│  │  │  ├─ Nav.js
│  │  │  └─ responsive.css
│  │  ├─ registerServiceWorker.js
│  │  ├─ Routes
│  │  │  ├─ Account
│  │  │  │  ├─ ChangePassword.js
│  │  │  │  └─ ChangePasswordForm.js
│  │  │  ├─ Helpers.js
│  │  │  ├─ Home
│  │  │  │  └─ Home.js
│  │  │  ├─ LogIn
│  │  │  │  ├─ LogIn.js
│  │  │  │  └─ LogInForm.js
│  │  │  ├─ Posts
│  │  │  │  ├─ Post.js
│  │  │  │  ├─ PostForm.js
│  │  │  │  └─ Posts.js
│  │  │  ├─ Reset
│  │  │  │  ├─ CheckReset.js
│  │  │  │  ├─ Reset.js
│  │  │  │  └─ ResetForm.js
│  │  │  ├─ Routes.js
│  │  │  ├─ SignUp
│  │  │  │  ├─ SignUp.js
│  │  │  │  └─ SignUpForm.js
│  │  │  └─ Verify
│  │  │     └─ Verify.js
│  │  ├─ semantic-ui
│  │  │  ├─ site
│  │  │  │  ├─ collections
│  │  │  │  │  ├─ breadcrumb.overrides
│  │  │  │  │  ├─ breadcrumb.variables
│  │  │  │  │  ├─ form.overrides
│  │  │  │  │  ├─ form.variables
│  │  │  │  │  ├─ grid.overrides
│  │  │  │  │  ├─ grid.variables
│  │  │  │  │  ├─ menu.overrides
│  │  │  │  │  ├─ menu.variables
│  │  │  │  │  ├─ message.overrides
│  │  │  │  │  ├─ message.variables
│  │  │  │  │  ├─ table.overrides
│  │  │  │  │  └─ table.variables
│  │  │  │  ├─ elements
│  │  │  │  │  ├─ button.overrides
│  │  │  │  │  ├─ button.variables
│  │  │  │  │  ├─ container.overrides
│  │  │  │  │  ├─ container.variables
│  │  │  │  │  ├─ divider.overrides
│  │  │  │  │  ├─ divider.variables
│  │  │  │  │  ├─ flag.overrides
│  │  │  │  │  ├─ flag.variables
│  │  │  │  │  ├─ header.overrides
│  │  │  │  │  ├─ header.variables
│  │  │  │  │  ├─ icon.overrides
│  │  │  │  │  ├─ icon.variables
│  │  │  │  │  ├─ image.overrides
│  │  │  │  │  ├─ image.variables
│  │  │  │  │  ├─ input.overrides
│  │  │  │  │  ├─ input.variables
│  │  │  │  │  ├─ label.overrides
│  │  │  │  │  ├─ label.variables
│  │  │  │  │  ├─ list.overrides
│  │  │  │  │  ├─ list.variables
│  │  │  │  │  ├─ loader.overrides
│  │  │  │  │  ├─ loader.variables
│  │  │  │  │  ├─ rail.overrides
│  │  │  │  │  ├─ rail.variables
│  │  │  │  │  ├─ reveal.overrides
│  │  │  │  │  ├─ reveal.variables
│  │  │  │  │  ├─ segment.overrides
│  │  │  │  │  ├─ segment.variables
│  │  │  │  │  ├─ step.overrides
│  │  │  │  │  └─ step.variables
│  │  │  │  ├─ globals
│  │  │  │  │  ├─ reset.overrides
│  │  │  │  │  ├─ reset.variables
│  │  │  │  │  ├─ site.overrides
│  │  │  │  │  └─ site.variables
│  │  │  │  ├─ modules
│  │  │  │  │  ├─ accordion.overrides
│  │  │  │  │  ├─ accordion.variables
│  │  │  │  │  ├─ chatroom.overrides
│  │  │  │  │  ├─ chatroom.variables
│  │  │  │  │  ├─ checkbox.overrides
│  │  │  │  │  ├─ checkbox.variables
│  │  │  │  │  ├─ dimmer.overrides
│  │  │  │  │  ├─ dimmer.variables
│  │  │  │  │  ├─ dropdown.overrides
│  │  │  │  │  ├─ dropdown.variables
│  │  │  │  │  ├─ embed.overrides
│  │  │  │  │  ├─ embed.variables
│  │  │  │  │  ├─ modal.overrides
│  │  │  │  │  ├─ modal.variables
│  │  │  │  │  ├─ nag.overrides
│  │  │  │  │  ├─ nag.variables
│  │  │  │  │  ├─ popup.overrides
│  │  │  │  │  ├─ popup.variables
│  │  │  │  │  ├─ progress.overrides
│  │  │  │  │  ├─ progress.variables
│  │  │  │  │  ├─ rating.overrides
│  │  │  │  │  ├─ rating.variables
│  │  │  │  │  ├─ search.overrides
│  │  │  │  │  ├─ search.variables
│  │  │  │  │  ├─ shape.overrides
│  │  │  │  │  ├─ shape.variables
│  │  │  │  │  ├─ sidebar.overrides
│  │  │  │  │  ├─ sidebar.variables
│  │  │  │  │  ├─ sticky.overrides
│  │  │  │  │  ├─ sticky.variables
│  │  │  │  │  ├─ tab.overrides
│  │  │  │  │  ├─ tab.variables
│  │  │  │  │  ├─ transition.overrides
│  │  │  │  │  └─ transition.variables
│  │  │  │  └─ views
│  │  │  │     ├─ ad.overrides
│  │  │  │     ├─ ad.variables
│  │  │  │     ├─ card.overrides
│  │  │  │     ├─ card.variables
│  │  │  │     ├─ comment.overrides
│  │  │  │     ├─ comment.variables
│  │  │  │     ├─ feed.overrides
│  │  │  │     ├─ feed.variables
│  │  │  │     ├─ item.overrides
│  │  │  │     ├─ item.variables
│  │  │  │     ├─ statistic.overrides
│  │  │  │     └─ statistic.variables
│  │  │  └─ theme.config
│  │  └─ Shared
│  │     ├─ Context.js
│  │     ├─ Hooks.js
│  │     ├─ Roles.js
│  │     └─ SimplePage.js
│  └─ webpack.config.js
├─ README.md
└─ __goInteract_basics

```


TODO

1. Добавить MongoDB (вместе или вместо PSQL, лучше вместе на всякий случай) в структуру и в API сервера
2. добавить реактовые версии старых страниц
3. добавить админку типа как-то?
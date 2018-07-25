# PNGR
dockerized (postgres + nginx + golang + react) starter kit

I've only implemented basic user management stuff in terms of actual features, but this scaffolding can be extended to serve a huge variety of purposes.

This project is meant to be a starting point. Feel free to create issues with suggestions, or pull requests for security improvements or developer ergonomics.

## Requirements
- Install docker && docker-compose

## Start the Dev Server
1) `sudo docker-compose up`
2) Visit https://localhost (and approve the self-signed cert)
3) Make changes to either front-end or back-end code, and watch it rebuild itself!

## Production Builds
*Warning: this code is pre-alpha quality! Run in production at your own risk*

- Generate a production container with `sudo docker build .` 
- You will need to setup nginx in production for SSL termination and port forwarding to `:5050` look at `nginx/nginx.prod.conf` for ideas on how to do this.

### Why serve front-end files with our API server?
It's setup this way so that we can block certain resources based on user authentication. Most notably to exclude unauthenticated users from download the admin javascript bundle.

This could be handled with something like JSON Web Tokens so that nginx can verify credentials without hitting the app.
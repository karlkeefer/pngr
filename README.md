# ofcourse
docker + golang + postgres + react

This project was built to showcase an app built using a few cool technologies.
I've only implemented basic user management stuff in terms of actual features, but this scaffolding can be extended to serve a huge variety of purposes.

I thought other people might be interested in the same stack, so it's open source.

## Getting started
- Install docker
- Install docker-compose
- Run the app using `sudo docker-compose up`

## Production builds
- Generate a production container with `sudo docker build .` 
- You will need to setup nginx in production for SSL termination and port forwarding to `:3000` look at `nginx/nginx.conf` for ideas on how to do this.

## Front-End Goodies
- React (un-ejected create-react-app)
- Semantic UI
- Unstated (alternative to redux)

## Back-End Goodies
- golang
- postgres
- docker

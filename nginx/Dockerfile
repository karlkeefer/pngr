FROM nginx:alpine as dev
COPY cert.pem /etc/nginx/conf.d
COPY key.pem /etc/nginx/conf.d
COPY nginx.conf /etc/nginx/nginx.conf

FROM nginx:alpine as prod
COPY cert.pem /etc/nginx/conf.d
COPY key.pem /etc/nginx/conf.d
COPY nginx.prod.conf /etc/nginx/nginx.conf
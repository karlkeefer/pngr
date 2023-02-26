FROM nginx:alpine as dev
COPY nginx.conf /etc/nginx/nginx.conf

FROM nginx:alpine as prod
COPY nginx.prod.conf /etc/nginx/nginx.conf

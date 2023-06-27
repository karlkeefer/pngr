FROM node:16-alpine as dev
WORKDIR /root/
CMD echo "getting latest npm modules" \
    && npm install \
    && echo "installation complete, starting" \
    && npm start

FROM node:16-alpine as build
WORKDIR /root
COPY ./package*.json ./
RUN npm ci --cache .npm
COPY . . 
RUN npm run build

# multistage build...
# this runs a lil nginx instance to serve the app
FROM nginx:alpine as prod
EXPOSE 3000
COPY --from=build /root/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf

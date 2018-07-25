# This dockerfile builds a dev version of the frontend
# It only really works as expected when run via docker-compose in the parent directory

FROM node:10.7-alpine
WORKDIR /root/
RUN npm install
CMD npm install && npm start
# develop stage
FROM node:14.0.0-alpine as develop-stage
WORKDIR /app
COPY ./package.json ./package.json[]
RUN yarn install --ignore-engines
COPY . .

# build stage
FROM develop-stage as build-stage
RUN yarn build

# production stage
FROM nginx:1.23.0-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
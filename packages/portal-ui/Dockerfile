# stage1 as builder
FROM node:16-alpine as builder

WORKDIR /portal-ui

# Copy the package.json and install dependencies
COPY package*.json ./
RUN yarn

# Copy rest of the files
COPY . .

# Build the project
RUN yarn build

FROM nginx:alpine as production-build
COPY ./.nginx/nginx.conf /etc/nginx/nginx.conf

## Remove default nginx index page
RUN rm -rf /usr/share/nginx/html/*

# Copy from the stahg 1
COPY --from=builder /portal-ui/dist /usr/share/nginx/html

EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]
# etapa 1: build del frontend
FROM node:20-alpine as build

WORKDIR /app
COPY . .

RUN npm install
RUN npm run build

# etapa 2: nginx para servirlo
FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

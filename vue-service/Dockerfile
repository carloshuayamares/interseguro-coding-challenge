FROM node:22-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
ARG VITE_GO_API_URL=http://localhost:3000
ARG VITE_GO_API_PATH=/api/v1/matrix/qr
ENV VITE_GO_API_URL=$VITE_GO_API_URL
ENV VITE_GO_API_PATH=$VITE_GO_API_PATH
RUN npm run build

FROM nginx:1.27-alpine AS production
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

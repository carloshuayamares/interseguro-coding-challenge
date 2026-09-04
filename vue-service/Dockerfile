FROM node:22-alpine AS build

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
ARG VITE_GO_API_URL=https://matrix-go-api-production.up.railway.app
ARG VITE_GO_API_PATH=/api/v1/matrix/qr
ARG VITE_JWT_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJtYXRyaXgtY2hhbGxlbmdlLWNsaWVudCIsImlhdCI6MTc4ODU0MjUwOSwiZXhwIjoxNzg5MTQ3MzA5fQ.MOBL6mSnLR4wPTZZQ0L0YmszXFSagxpjRgSZry4HKt8
ENV VITE_GO_API_URL=$VITE_GO_API_URL
ENV VITE_GO_API_PATH=$VITE_GO_API_PATH
ENV VITE_JWT_TOKEN=$VITE_JWT_TOKEN
RUN npm run build

FROM nginx:1.27-alpine AS production

COPY nginx.conf /etc/nginx/templates/default.conf.template
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

# Matrix QR Service

Microservicio REST en Go/Fiber para validar matrices, calcular su factorizacion QR y enviar `q` y `r` a un servicio Node.js.

## Configuracion

- `PORT`: puerto HTTP, por defecto `8080`.
- `JWT_SECRET`: secreto HMAC para validar tokens JWT.
- `NODE_SERVICE_URL`: URL del endpoint Node.js, por defecto `http://node-service:3000/analyze`.

## Ejecutar

```bash
go mod tidy
go test ./...
go run .
```

## Peticion

```bash
curl -X POST http://localhost:8080/api/v1/matrix/process \
  -H "Authorization: Bearer <jwt>" \
  -H "Content-Type: application/json" \
  -d '[[1,2],[3,4]]'
```

El JWT se reenvia al servicio Node.js como cabecera `Authorization`.

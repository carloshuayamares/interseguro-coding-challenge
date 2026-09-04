# Node Stats Service

Microservicio Express/TypeScript que valida matrices, calcula estadisticas agregadas y comprueba si alguna matriz es diagonal.

## Variables de entorno

- `PORT`: puerto HTTP, por defecto `4000`.
- `JWT_SECRET`: secreto usado para validar el JWT HMAC.

## Ejecucion local

```bash
npm install
npm test
npm run build
npm start
```

## Endpoint

`POST /api/v1/stats`

Headers:

```text
Authorization: Bearer <jwt>
Content-Type: application/json
```

Body:

```json
{"matrices":[[[1,2],[3,4]],[[5,6],[7,8]]]}
```

Respuesta:

```json
{"min":1,"max":8,"average":4.5,"sum":36,"hasDiagonalMatrix":false}
```

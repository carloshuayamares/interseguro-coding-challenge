# Matrix QR Challenge

Aplicacion compuesta por tres microservicios:

```text
Vue 3 -> Go/Fiber -> Node.js/Express
```

## Servicios

- `vue-service`: permite editar una matriz y muestra la matriz original, `Q`, `R` y las estadisticas.
- `go-matrix-service`: valida la matriz, calcula la factorizacion QR y coordina la llamada a Node.
- `node-stats-service`: calcula minimo, maximo, promedio, suma y detecta matrices diagonales.

## Ejecutar con Docker

1. Genera un secreto para la prueba y un JWT de 7 dias de expiracion. El secreto debe ser el mismo para Go y Node. En Windows PowerShell:

```powershell
$env:JWT_SECRET = "local-challenge-secret"
Push-Location node-stats-service
npm run token
Pop-Location
```

Usa la ultima linea impresa por el comando como valor de `VITE_JWT_TOKEN`.

2. Crea un archivo `.env` en la raiz con el mismo secreto y el token generado:

```dotenv
JWT_SECRET=local-challenge-secret
VITE_JWT_TOKEN=pega-aqui-el-token-generado
```

3. Levanta todos los servicios:

```powershell
docker compose up --build
```

URLs:

- Frontend: `http://localhost:5173`
- Go API: `http://localhost:3000`
- Node API: `http://localhost:4000`

## Token JWT

Para esta prueba se usa un JWT firmado con `HS256`, compartiendo `JWT_SECRET` entre Go y Node. El comando `npm run token` genera un token con expiracion de `7d`, suficiente para una evaluacion o demostracion sin renovaciones manuales.

Las variables tienen responsabilidades distintas:

- `JWT_SECRET`: variable de runtime para Go y Node. Ambos servicios deben usar el mismo valor para validar la firma.
- `VITE_JWT_TOKEN`: argumento de build para Vue. Se inserta en el bundle estatico y el navegador lo envia en `Authorization`.

Por eso el token no se genera en el `Dockerfile` de Node: ese contenedor no crea tokens, solo los valida. Tampoco se configura `VITE_JWT_TOKEN` como `environment` del contenedor frontend, porque Nginx sirve un bundle ya compilado. El token se genera una vez antes de `docker compose up --build` y Compose lo pasa al build de Vue.

No se necesita un endpoint de renovacion para este challenge. En un entorno real el frontend no deberia recibir un secreto compartido ni un JWT permanente: se usaria un proveedor de identidad, tokens de corta duracion y refresh tokens. Como `VITE_JWT_TOKEN` queda incluido en el bundle del navegador, debe considerarse una credencial de prueba, no un secreto de produccion.

## Endpoints

Todos los endpoints de negocio requieren:

```text
Authorization: Bearer <JWT>
Content-Type: application/json
```

### Go: QR

`POST http://localhost:3000/api/v1/matrix/qr`

```json
[[1, 2], [3, 4], [5, 6]]
```

Devuelve `q`, `r` y `analysis`. Go envia `Q` y `R` a Node para obtener las estadisticas.

### Go: rotacion

`POST http://localhost:3000/api/v1/matrix/rotate`

Devuelve la matriz rotada 90 grados y sus estadisticas.

### Node: estadisticas

`POST http://localhost:4000/api/v1/stats`

```json
{"matrices": [[[1, 0], [0, 1]]]}
```

## Desarrollo local

Go:

```powershell
cd go-matrix-service
go test ./...
go run .
```

Node:

```powershell
cd node-stats-service
npm install
npm test
npm run build
```

Vue:

```powershell
cd vue-service
npm install
npm run dev
```
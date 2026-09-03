import 'dotenv/config';
import express from 'express';
import { postStats } from './controllers/stats.controller';
import { jwtMiddleware } from './middlewares/jwt.middleware';

const app = express();
const port = Number(process.env.PORT ?? 3000);
const jwtSecret = process.env.JWT_SECRET ?? 'change-me-in-production';

app.use(express.json({ limit: '1mb' }));
app.get('/health', (_request, response) => response.json({ status: 'ok' }));
app.post('/api/v1/stats', jwtMiddleware(jwtSecret), postStats);

app.listen(port, () => {
  console.log(`node stats service listening on :${port}`);
});

export { app };

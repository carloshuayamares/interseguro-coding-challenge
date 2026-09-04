import { NextFunction, Request, Response } from 'express';
import jwt from 'jsonwebtoken';

export function jwtMiddleware(secret: string) {
  return (request: Request, response: Response, next: NextFunction): void => {
    const authorization = request.header('authorization');
    const [scheme, token] = authorization?.split(' ') ?? [];

    if (scheme?.toLowerCase() !== 'bearer' || !token) {
      response.status(401).json({ error: 'missing or invalid bearer token' });
      return;
    }

    try {
      jwt.verify(token, secret, { algorithms: ['HS256'] });
      next();
    } catch {
      response.status(401).json({ error: 'invalid bearer token' });
    }
  };
}

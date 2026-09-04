const jwt = require('jsonwebtoken');

const secret = process.env.JWT_SECRET || 'change-me-in-production';
const token = jwt.sign(
  { sub: process.env.JWT_SUBJECT || 'matrix-challenge-client' },
  secret,
  { algorithm: 'HS256', expiresIn: '7d' }
);

console.log(token);

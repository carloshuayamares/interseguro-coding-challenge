import axios from 'axios';
import type { Matrix, QRResponse } from '../types/matrix';

const api = axios.create({
  baseURL: import.meta.env.VITE_GO_API_URL,
  headers: { 'Content-Type': 'application/json' },
  timeout: 15000
});

export async function calculateQR(matrix: Matrix): Promise<QRResponse> {
  const token = import.meta.env.VITE_JWT_TOKEN?.trim();
  const response = await api.post<QRResponse>(
    import.meta.env.VITE_GO_API_PATH || '/api/v1/matrix/qr',
    matrix,
    token ? { headers: { Authorization: `Bearer ${token}` } } : undefined
  );
  return response.data;
}

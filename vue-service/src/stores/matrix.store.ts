import { defineStore } from 'pinia';
import axios from 'axios';
import { calculateQR } from '../services/goApi';
import type { Matrix, QRResponse, Statistics } from '../types/matrix';

interface MatrixState {
  matrix: Matrix;
  qMatrix: Matrix;
  rMatrix: Matrix;
  statistics: Statistics | null;
  loading: boolean;
  error: string | null;
}

function getErrorMessage(error: unknown): string {
  if (axios.isAxiosError(error)) {
    if (!error.response) return 'Go API is unavailable.';
    if (error.response.status >= 500) return 'Unable to process matrix. Please try again.';
    return 'Unable to process matrix. Please verify the matrix dimensions.';
  }
  return 'Unable to process matrix. Please try again.';
}

export const useMatrixStore = defineStore('matrix', {
  state: (): MatrixState => ({
    matrix: [],
    qMatrix: [],
    rMatrix: [],
    statistics: null,
    loading: false,
    error: null
  }),
  actions: {
    async calculate(matrix: Matrix): Promise<void> {
      if (this.loading) return;
      this.loading = true;
      this.error = null;
      this.matrix = matrix.map((row) => [...row]);
      try {
        const result: QRResponse = await calculateQR(matrix);
        this.qMatrix = result.q;
        this.rMatrix = result.r;
        this.statistics = result.analysis;
      } catch (error: unknown) {
        this.error = getErrorMessage(error);
      } finally {
        this.loading = false;
      }
    },
    clearResult(): void {
      this.matrix = [];
      this.qMatrix = [];
      this.rMatrix = [];
      this.statistics = null;
    },
    clearError(): void {
      this.error = null;
    }
  }
});

export type NumericMatrix = number[][];

export interface Statistics {
  min: number;
  max: number;
  average: number;
  sum: number;
  hasDiagonalMatrix: boolean;
}

export function validateMatrices(matrices: unknown): matrices is NumericMatrix[] {
  if (!Array.isArray(matrices) || matrices.length === 0) {
    return false;
  }

  return matrices.every((matrix) => {
    if (!Array.isArray(matrix) || matrix.length === 0) {
      return false;
    }
    const columns = Array.isArray(matrix[0]) ? matrix[0].length : 0;
    if (columns === 0) {
      return false;
    }
    return matrix.every(
      (row) =>
        Array.isArray(row) &&
        row.length === columns &&
        row.every((value) => typeof value === 'number' && Number.isFinite(value))
    );
  });
}

export function calculateStatistics(matrices: NumericMatrix[]): Statistics {
  if (!validateMatrices(matrices)) {
    throw new Error('matrices must be a non-empty array of rectangular numeric matrices');
  }

  const values = matrices.flat(2);
  const sum = values.reduce((total, value) => total + value, 0);

  const response: Statistics = {
    min: Math.min(...values),
    max: Math.max(...values),
    average: sum / values.length,
    sum,
    hasDiagonalMatrix: matrices.some(isDiagonalMatrix)
  };
  console.log('Calculated statistics:', response);
  return response;
}

export function isDiagonalMatrix(matrix: NumericMatrix): boolean {
  if (matrix.length === 0 || matrix.some((row) => row.length !== matrix.length)) {
    return false;
  }

  return matrix.every((row, rowIndex) =>
    row.every((value, columnIndex) => rowIndex === columnIndex || value === 0)
  );
}

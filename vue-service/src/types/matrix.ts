export type Matrix = number[][];

export interface QRResponse {
  q: Matrix;
  r: Matrix;
  analysis: Statistics;
}

export interface Statistics {
  min: number;
  max: number;
  average: number;
  sum: number;
  hasDiagonalMatrix: boolean;
}

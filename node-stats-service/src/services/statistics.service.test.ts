import {
  calculateStatistics,
  isDiagonalMatrix,
  validateMatrices
} from './statistics.service';
import { describe, expect, test } from '@jest/globals';

describe('validateMatrices', () => {
  test('accepts non-empty rectangular numeric matrices', () => {
    expect(validateMatrices([[[1, 2], [3, 4]], [[5, 6]]])).toBe(true);
  });

  test('rejects empty, ragged, and non-numeric matrices', () => {
    expect(validateMatrices([])).toBe(false);
    expect(validateMatrices([[[1], [2, 3]]])).toBe(false);
    expect(validateMatrices([[[1, '2' as unknown as number]]])).toBe(false);
  });
});

describe('calculateStatistics', () => {
  test('calculates min, max, average, and sum across all matrices', () => {
    expect(calculateStatistics([[[1, 2], [3, 4]], [[5, 6], [7, 8]]])).toEqual({
      min: 1,
      max: 8,
      average: 4.5,
      sum: 36,
      hasDiagonalMatrix: false
    });
  });

  test('reports a diagonal matrix when any matrix is diagonal', () => {
    expect(calculateStatistics([[[1, 0], [0, 2]], [[1, 1], [0, 1]]]).hasDiagonalMatrix).toBe(true);
  });
});

describe('isDiagonalMatrix', () => {
  test('accepts square matrices with zero off the main diagonal', () => {
    expect(isDiagonalMatrix([[2, 0], [0, -1]])).toBe(true);
  });

  test('rejects non-square or non-diagonal matrices', () => {
    expect(isDiagonalMatrix([[1, 0, 0], [0, 1, 0]])).toBe(false);
    expect(isDiagonalMatrix([[1, 2], [0, 1]])).toBe(false);
  });
});

import { Request, Response } from 'express';
import { calculateStatistics, NumericMatrix, validateMatrices } from '../services/statistics.service';

interface StatsRequestBody {
  matrices?: unknown;
}

export function postStats(request: Request, response: Response): void {
  const { matrices } = (request.body ?? {}) as StatsRequestBody;

  if (!validateMatrices(matrices)) {
    response.status(400).json({
      error: 'matrices must be a non-empty array of rectangular numeric matrices'
    });
    return;
  }

  const statistics = calculateStatistics(matrices as NumericMatrix[]);
  response.status(200).json(statistics);
}

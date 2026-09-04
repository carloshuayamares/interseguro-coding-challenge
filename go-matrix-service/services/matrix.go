package services

import (
	"context"
	"errors"
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

var (
	ErrEmptyMatrix    = errors.New("matrix must not be empty")
	ErrEmptyRow       = errors.New("matrix rows must not be empty")
	ErrNonRectangular = errors.New("matrix must be rectangular")
	ErrNonFiniteValue = errors.New("matrix values must be finite")
	ErrSingularQR     = errors.New("matrix cannot be factorized into QR")
)

type Matrix [][]float64

type QRResult struct {
	Q [][]float64 `json:"q"`
	R [][]float64 `json:"r"`
}

type Analyzer interface {
	Analyze(context.Context, QRResult, string) (interface{}, error)
}

type MatrixAnalyzer interface {
	AnalyzeMatrices(context.Context, []Matrix, string) (interface{}, error)
}

type MatrixService struct {
	analyzer Analyzer
}

func NewMatrixService(analyzer Analyzer) *MatrixService {
	return &MatrixService{analyzer: analyzer}
}

func (s *MatrixService) Process(ctx context.Context, input Matrix, token string) (QRResult, interface{}, error) {

	if err := ValidateMatrix(input); err != nil {
		return QRResult{}, nil, err
	}

	result, err := FactorizeQR(input)
	if err != nil {
		return QRResult{}, nil, err
	}
	analysis, err := s.analyzer.Analyze(ctx, result, token)
	if err != nil {
		return QRResult{}, nil, fmt.Errorf("node analysis: %w", err)
	}
	return result, analysis, nil
}

func (s *MatrixService) Rotate(ctx context.Context, input Matrix, token string) (Matrix, interface{}, error) {
	if err := ValidateMatrix(input); err != nil {
		return nil, nil, err
	}

	rotated := RotateClockwise(input)
	analyzer, ok := s.analyzer.(MatrixAnalyzer)
	if !ok {
		return nil, nil, errors.New("configured analyzer does not support matrix analysis")
	}
	analysis, err := analyzer.AnalyzeMatrices(ctx, []Matrix{rotated}, token)
	if err != nil {
		return nil, nil, fmt.Errorf("node analysis: %w", err)
	}
	return rotated, analysis, nil
}

func RotateClockwise(input Matrix) Matrix {
	rows, columns := len(input), len(input[0])
	rotated := make(Matrix, columns)
	for row := range rotated {
		rotated[row] = make([]float64, rows)
		for column := 0; column < rows; column++ {
			rotated[row][column] = input[rows-1-column][row]
		}
	}
	return rotated
}

func ValidateMatrix(input Matrix) error {
	if len(input) == 0 {
		return ErrEmptyMatrix
	}
	if len(input[0]) == 0 {
		return ErrEmptyRow
	}
	columns := len(input[0])
	for _, row := range input {
		if len(row) == 0 {
			return ErrEmptyRow
		}
		if len(row) != columns {
			return ErrNonRectangular
		}
		for _, value := range row {
			if math.IsNaN(value) || math.IsInf(value, 0) {
				return ErrNonFiniteValue
			}
		}
	}
	return nil
}

func FactorizeQR(input Matrix) (QRResult, error) {
	if err := ValidateMatrix(input); err != nil {
		return QRResult{}, err
	}
	rows, columns := len(input), len(input[0])
	values := make([]float64, 0, rows*columns)
	for _, row := range input {
		values = append(values, row...)
	}

	factorization := mat.QR{}
	factorization.Factorize(mat.NewDense(rows, columns, values))
	qFull := mat.NewDense(rows, rows, nil)
	factorization.QTo(qFull)
	q := mat.NewDense(rows, min(rows, columns), nil)
	for row := 0; row < rows; row++ {
		for column := 0; column < min(rows, columns); column++ {
			q.Set(row, column, qFull.At(row, column))
		}
	}
	rFull := mat.NewDense(rows, columns, nil)
	factorization.RTo(rFull)
	r := mat.NewDense(min(rows, columns), columns, nil)
	for row := 0; row < min(rows, columns); row++ {
		for column := 0; column < columns; column++ {
			r.Set(row, column, rFull.At(row, column))
		}
	}
	return QRResult{Q: denseToSlices(q), R: denseToSlices(r)}, nil
}

func denseToSlices(matrix *mat.Dense) [][]float64 {
	rows, columns := matrix.Dims()
	result := make([][]float64, rows)
	for row := range result {
		result[row] = make([]float64, columns)
		for column := range result[row] {
			result[row][column] = matrix.At(row, column)
		}
	}
	return result
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}

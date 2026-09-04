package services

import (
	"context"
	"math"
	"testing"
)

type analyzerStub struct{}

func (analyzerStub) Analyze(context.Context, QRResult, string) (interface{}, error) {
	return map[string]string{"status": "ok"}, nil
}

func (analyzerStub) AnalyzeMatrices(context.Context, []Matrix, string) (interface{}, error) {
	return map[string]string{"status": "ok"}, nil
}

func TestValidateMatrix(t *testing.T) {
	tests := []struct {
		name  string
		input Matrix
		want  error
	}{
		{"empty", Matrix{}, ErrEmptyMatrix},
		{"empty row", Matrix{{}}, ErrEmptyRow},
		{"non rectangular", Matrix{{1, 2}, {3}}, ErrNonRectangular},
		{"nan", Matrix{{math.NaN()}}, ErrNonFiniteValue},
		{"valid", Matrix{{1, 2}, {3, 4}}, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := ValidateMatrix(test.input); got != test.want {
				t.Fatalf("ValidateMatrix() error = %v, want %v", got, test.want)
			}
		})
	}
}

func TestFactorizeQR(t *testing.T) {
	input := Matrix{{1, 2}, {3, 4}, {5, 6}}
	result, err := FactorizeQR(input)
	if err != nil {
		t.Fatalf("FactorizeQR() error = %v", err)
	}
	if len(result.Q) != 3 || len(result.Q[0]) != 2 || len(result.R) != 2 || len(result.R[0]) != 2 {
		t.Fatalf("unexpected dimensions: Q=%dx%d R=%dx%d", len(result.Q), len(result.Q[0]), len(result.R), len(result.R[0]))
	}
	for row := range input {
		for column := range input[0] {
			var reconstructed float64
			for index := range result.R {
				reconstructed += result.Q[row][index] * result.R[index][column]
			}
			if math.Abs(reconstructed-input[row][column]) > 1e-10 {
				t.Fatalf("A != QR at (%d,%d): got %v, want %v", row, column, reconstructed, input[row][column])
			}
		}
	}
}

func TestMatrixServiceProcessCallsAnalyzer(t *testing.T) {
	service := NewMatrixService(analyzerStub{})
	_, analysis, err := service.Process(context.Background(), Matrix{{1}}, "token")
	if err != nil {
		t.Fatalf("Process() error = %v", err)
	}
	if analysis == nil {
		t.Fatal("Process() returned nil analysis")
	}
}

func TestRotateClockwise(t *testing.T) {
	input := Matrix{{1, 2, 3}, {4, 5, 6}}
	want := Matrix{{4, 1}, {5, 2}, {6, 3}}
	got := RotateClockwise(input)
	if len(got) != len(want) || len(got[0]) != len(want[0]) {
		t.Fatalf("RotateClockwise() dimensions = %dx%d, want %dx%d", len(got), len(got[0]), len(want), len(want[0]))
	}
	for row := range want {
		for column := range want[row] {
			if got[row][column] != want[row][column] {
				t.Fatalf("RotateClockwise()[%d][%d] = %v, want %v", row, column, got[row][column], want[row][column])
			}
		}
	}
}

func TestMatrixServiceRotateCallsAnalyzer(t *testing.T) {
	service := NewMatrixService(analyzerStub{})
	rotated, analysis, err := service.Rotate(context.Background(), Matrix{{1, 2}, {3, 4}}, "token")
	if err != nil {
		t.Fatalf("Rotate() error = %v", err)
	}
	if rotated[0][0] != 3 || rotated[1][1] != 2 || analysis == nil {
		t.Fatalf("Rotate() returned rotated=%v, analysis=%v", rotated, analysis)
	}
}

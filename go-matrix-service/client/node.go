package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/interseguro/matrix-service/services"
)

type HTTPClient struct {
	Client *http.Client
	URL    string
}

type analysisRequest struct {
	Matrices []services.Matrix `json:"matrices"`
}

func (h *HTTPClient) Analyze(ctx context.Context, qr services.QRResult, token string) (interface{}, error) {
	return h.AnalyzeMatrices(ctx, []services.Matrix{qr.Q, qr.R}, token)
}

func (h *HTTPClient) AnalyzeMatrices(ctx context.Context, matrices []services.Matrix, token string) (interface{}, error) {
	payload, err := json.Marshal(analysisRequest{Matrices: matrices})
	if err != nil {
		return nil, fmt.Errorf("marshal analysis payload: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, h.URL, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("create analysis request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := h.Client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("send analysis request: %w", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read analysis response: %w", err)
	}
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("analysis service returned status %d: %s", response.StatusCode, body)
	}
	if len(body) == 0 {
		return map[string]interface{}{}, nil
	}
	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("decode analysis response: %w", err)
	}
	return result, nil
}

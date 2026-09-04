package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/interseguro/matrix-service/services"
)

type MatrixHandler struct {
	service *services.MatrixService
}

func NewMatrixHandler(service *services.MatrixService) *MatrixHandler {
	return &MatrixHandler{service: service}
}

func (h *MatrixHandler) Process(c *fiber.Ctx) error {
	// escribir en consola el body de la request
	log.Printf("Received matrix: %s", c.Body())

	var matrix services.Matrix
	if err := json.Unmarshal(c.Body(), &matrix); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON body"})
	}

	result, analysis, err := h.service.Process(c.UserContext(), matrix, c.Locals("jwt").(string))
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrEmptyMatrix) || errors.Is(err, services.ErrEmptyRow) || errors.Is(err, services.ErrNonRectangular) || errors.Is(err, services.ErrNonFiniteValue) {
			status = http.StatusBadRequest
		}
		return c.Status(status).JSON(fiber.Map{"error": err.Error()})
	}
	log.Printf("Processed matrix: %s, result: %v, analysis: %v", c.Body(), result, analysis)

	return c.JSON(fiber.Map{"q": result.Q, "r": result.R, "analysis": analysis})
}

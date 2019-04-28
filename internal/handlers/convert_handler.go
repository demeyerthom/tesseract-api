package handlers

import (
	"context"
	"github.com/demeyerthom/tesseract-api/internal"
	"github.com/demeyerthom/tesseract-api/pkg/tesseract"
	"github.com/go-openapi/runtime/middleware"
)

type ConvertHandler struct {
	tesseractClient *tesseract.Client
}

func NewConvertHandler(client *tesseract.Client) *ConvertHandler {
	return &ConvertHandler{tesseractClient: client}
}

func (h *ConvertHandler) Handle(params internal.ConvertParams) middleware.Responder {
	bytes, _ := h.tesseractClient.Convert(context.Background(), params.File)

	text := string(bytes)

	parsedImage := internal.ParsedImage{
		Text: &text,
	}

	response := internal.NewConvertOK()

	response.WithPayload(&parsedImage)

	return response
}

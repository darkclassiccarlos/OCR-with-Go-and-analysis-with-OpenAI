package openai

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

// AnalyzeText envía el texto extraído a la API de OpenAI y obtiene una respuesta
func AnalyzeText(ctx context.Context, apiKey string, extractedText string) (string, error) {
	client := openai.NewClient(apiKey)

	// Crear la solicitud para OpenAI
	response, err := client.CreateCompletion(ctx, openai.CompletionRequest{
		Model:     "text-davinci-003",                                   // Usar el modelo que necesites
		Prompt:    fmt.Sprintf("Analiza este texto: %s", extractedText), // Prompt personalizado
		MaxTokens: 100,                                                  // Limitar la longitud de la respuesta
	})
	if err != nil {
		return "", err
	}

	return response.Choices[0].Text, nil
}

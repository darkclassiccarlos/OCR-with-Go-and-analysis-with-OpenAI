package api

import (
	"context"
	"net/http"
	"os"

	"go-image-text/internal/ocr"
	"go-image-text/internal/openai"

	"github.com/gin-gonic/gin"
)

func AnalyzeImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image file provided"})
		return
	}

	c.SaveUploadedFile(file, "uploaded_image.png")

	// Extraer texto usando Tesseract
	text, err := ocr.ExtractTextFromImage("uploaded_image.png")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing image"})
		return
	}

	// Enviar texto a OpenAI para su análisis
	apiKey := os.Getenv("OPENAI_API_KEY")
	ctx := context.Background()
	response, err := openai.AnalyzeText(ctx, apiKey, text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error analyzing text"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"text":     text,
		"response": response,
	})
}

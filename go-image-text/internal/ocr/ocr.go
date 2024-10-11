package ocr

import (
	"log"

	"github.com/otiai10/gosseract/v2"
)

// ExtractTextFromImage usa Tesseract para extraer texto de una imagen
func ExtractTextFromImage(imagePath string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage(imagePath)
	text, err := client.Text()
	if err != nil {
		log.Printf("Error al extraer texto de la imagen: %v", err)
		return "", err
	}

	return text, nil
}

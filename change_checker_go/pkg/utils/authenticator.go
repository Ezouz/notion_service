package utils

import (
	"log"
	"os"

	"github.com/go-openapi/errors"
	"gitlab.42paris.fr/notion_service/pkg/models"
)

func ValidateHeader(bearerHeader string) (interface{}, error) {
	log.Printf("VALIDATEHEADER: %s // %s", bearerHeader, os.Getenv("API_SECRET_ANSIBLE"))

	if bearerHeader == os.Getenv("API_SECRET_ANSIBLE") {
		return models.Principal(bearerHeader), nil
		// return bearerHeader, nil
	}

	return nil, errors.New(401, "Incorrect API key")
}

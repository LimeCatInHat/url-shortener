package app

import (
	"errors"
	"fmt"

	"github.com/LimeCatInHat/url-shortener/internal/config"
	"github.com/LimeCatInHat/url-shortener/internal/utils"
)

const keyGenerationAttemptsLimit = 5

type URLStogare interface {
	HasKey(key string) bool
	GetFullURL(key string) (string, error)
	GetShortKey(fullURL string) (string, error)
	SaveURLByShortKey(key string, value string)
}

func ShortenURL(url []byte, stor URLStogare) (string, error) {
	value, err := stor.GetShortKey(string(url))
	if err == nil {
		return getShortenURL(value), nil
	}

	key, err := generateNewKey(stor)
	if err != nil {
		return "", err
	}

	stor.SaveURLByShortKey(key, string(url))
	return getShortenURL(key), nil
}

func GetFullURL(key []byte, stor URLStogare) (string, error) {
	url, err := stor.GetFullURL(string(key))
	if err != nil {
		return "", fmt.Errorf("getting full url failed: %w", err)
	}
	return url, nil
}

func getShortenURL(key string) string {
	baseURL := config.GetConfiguration().ShortenLinksBaseURL
	return baseURL + key
}

func generateNewKey(stor URLStogare) (string, error) {
	for i := keyGenerationAttemptsLimit; i > 0; i-- {
		key := utils.GenerateRandomKey()
		if !stor.HasKey(key) {
			return key, nil
		}
	}
	return "", errors.New("max attempts count to generate new key exceeded")
}

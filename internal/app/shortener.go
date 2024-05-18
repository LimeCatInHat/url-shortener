package app

import (
	"errors"
	"fmt"

	"github.com/LimeCatInHat/url-shortener/internal/config"
	"github.com/LimeCatInHat/url-shortener/internal/utils"
)

const keyGenerationAttemptsLimit = 5

type URLStorage interface {
	GetFullURL(shortURL string) (string, error)
	GetShortURL(fullURL string) (string, error)
	SaveURL(fullURL string, shortURL string) error
}

func ShortenURL(url []byte, stor URLStorage) (string, error) {
	value, err := stor.GetShortURL(string(url))
	if err == nil {
		return getShortenURL(value), nil
	}

	key, err := generateShortKey(stor, url)
	if err != nil {
		return "", err
	}
	return getShortenURL(key), nil
}

func GetFullURL(key []byte, stor URLStorage) (string, error) {
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

func generateShortKey(stor URLStorage, url []byte) (string, error) {
	for range keyGenerationAttemptsLimit {
		shortURL := utils.GenerateRandomKey()
		err := stor.SaveURL(string(url), shortURL)
		if err == nil {
			return shortURL, nil
		}
	}
	return "", errors.New("max attempts count to generate new key exceeded")
}

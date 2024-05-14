package app

import (
	"errors"

	"github.com/LimeCatInHat/url-shortener/internal/config"
	"github.com/LimeCatInHat/url-shortener/internal/storage"
	"github.com/LimeCatInHat/url-shortener/internal/utils"
)

const keyGenerationAttemptsLimit = 5

func ShortenURL(url []byte, storage storage.URLStogare) (string, error) {
	isFound, value := storage.TryGetShortKey(string(url))
	if isFound {
		return getShortenURL(value), nil
	}

	key, err := generateNewKey(storage)
	if err != nil {
		return "", err
	}

	storage.SaveURLByShortKey(key, string(url))
	return getShortenURL(key), nil
}

func TryGetFullURL(key []byte, storage storage.URLStogare) (bool, string) {
	return storage.TryGetFullURL(string(key))
}

func getShortenURL(key string) string {
	baseURL := config.GetConfiguration().ShortenLinksBaseURL
	return baseURL + key
}

func generateNewKey(storage storage.URLStogare) (string, error) {
	for i := keyGenerationAttemptsLimit; i > 0; i-- {
		key := utils.GenerateRandomKey()
		if storage.HasKey(key) {
			continue
		}
		return key, nil
	}
	return "", errors.New("max attempts count to generate new key exceeded")
}

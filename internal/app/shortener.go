package app

import (
	"github.com/LimeCatInHat/url-shortener/internal/configuration"
	"github.com/LimeCatInHat/url-shortener/internal/storage"
	"github.com/LimeCatInHat/url-shortener/internal/utils"
)

var appStorage storage.IURLStogare = storage.AppMemoryStorage

func ConfigureStorage(storage storage.IURLStogare) {
	appStorage = storage
}

func ShortenURL(url []byte) string {
	isFound, value := appStorage.TryGetShortKey(string(url))
	if isFound {
		return getShortenURL(value)
	}
	key := utils.GenerateKey(url)

	appStorage.SaveURLByShortKey(key, string(url))

	return getShortenURL(key)
}

func TryGetFullURL(key []byte) (bool, string) {
	return appStorage.TryGetFullURL(string(key))
}

func getShortenURL(key string) string {
	return configuration.ShortenLinksBaseURL + key
}

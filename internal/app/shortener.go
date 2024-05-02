package app

import (
	"github.com/LimeCatInHat/url-shortener/internal/configuration"
	"github.com/LimeCatInHat/url-shortener/internal/storage"
	"github.com/LimeCatInHat/url-shortener/internal/utils"
)

var appStorage storage.IUrlStogare = storage.AppMemoryStorage

func ShortenUrl(url []byte) string {
	isFound, value := appStorage.TryGetShortKey(string(url))
	if isFound {
		return getShortenUrl(value)
	}
	key := utils.GenerateKey(url)

	appStorage.SaveUrlByShortKey(key, string(url))

	return getShortenUrl(key)
}

func TryGetFullUrl(key []byte) (bool, string) {
	return appStorage.TryGetFullUrl(string(key))
}

func getShortenUrl(key string) string {
	return configuration.ShortenLinksBaseUrl + key
}

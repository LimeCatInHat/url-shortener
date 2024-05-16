package storage

import (
	"fmt"
)

type MemoryStorage struct {
	urls map[string]string
}

func GetStorage() MemoryStorage {
	return MemoryStorage{urls: make(map[string]string)}
}

func (stor MemoryStorage) HasKey(key string) bool {
	_, found := stor.urls[key]
	return found
}

func (stor MemoryStorage) GetFullURL(key string) (string, error) {
	value, found := stor.urls[key]
	if found {
		return value, nil
	}
	return "", fmt.Errorf("attempt to get full URL for short link '%q' failed", key)
}

func (stor MemoryStorage) GetShortKey(fullURL string) (string, error) {
	for key, value := range stor.urls {
		if value == fullURL {
			return key, nil
		}
	}
	return "", fmt.Errorf("attempt to get shorten link for '%q' failed", fullURL)
}

func (stor MemoryStorage) SaveURLByShortKey(key string, value string) {
	stor.urls[key] = value
}

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

func (stor MemoryStorage) GetFullURL(shortURL string) (string, error) {
	value, found := stor.urls[shortURL]
	if found {
		return value, nil
	}
	return "", fmt.Errorf("attempt to get full URL for short link '%q' failed", shortURL)
}

func (stor MemoryStorage) GetShortURL(fullURL string) (string, error) {
	for key, value := range stor.urls {
		if value == fullURL {
			return key, nil
		}
	}
	return "", fmt.Errorf("attempt to get shorten link for '%q' failed", fullURL)
}

func (stor MemoryStorage) SaveURL(fullURL string, shortURL string) error {
	if _, exists := stor.urls[shortURL]; exists {
		return fmt.Errorf("key '%q' already exists", shortURL)
	}
	stor.urls[shortURL] = fullURL
	return nil
}

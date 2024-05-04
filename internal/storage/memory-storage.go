package storage

type MemoryStorage struct {
	urls map[string]string
}

var AppMemoryStorage IURLStogare = MemoryStorage{urls: make(map[string]string)}

func (storage MemoryStorage) TryGetFullURL(key string) (bool, string) {
	value, found := storage.urls[key]
	return found, value
}

func (storage MemoryStorage) TryGetShortKey(fullURL string) (bool, string) {
	for key, value := range storage.urls {
		if value == fullURL {
			return true, key
		}
	}
	return false, ""
}

func (storage MemoryStorage) SaveURLByShortKey(key string, value string) {
	storage.urls[key] = value
}

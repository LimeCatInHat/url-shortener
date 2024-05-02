package storage

type MemoryStorage struct {
	urls map[string]string
}

var AppMemoryStorage IUrlStogare = MemoryStorage{urls: make(map[string]string)}

func (storage MemoryStorage) TryGetFullUrl(key string) (isFound bool, value string) {
	value = storage.urls[key]
	return value != "", value
}

func (storage MemoryStorage) TryGetShortKey(fullUrl string) (isFound bool, value string) {
	for key, value := range storage.urls {
		if value == fullUrl {
			return true, key
		}
	}
	return false, ""
}

func (storage MemoryStorage) SaveUrlByShortKey(key string, value string) {
	storage.urls[key] = value
}

package storage

type MemoryStorage struct {
	urls map[string]string
}

type URLStogare interface {
	HasKey(key string) bool
	TryGetFullURL(key string) (isSucceed bool, value string)
	TryGetShortKey(fullURL string) (isSucceed bool, value string)
	SaveURLByShortKey(key string, value string)
}

func (stor MemoryStorage) HasKey(key string) bool {
	_, found := stor.urls[key]
	return found
}

func GetStorage() MemoryStorage {
	return MemoryStorage{urls: make(map[string]string)}
}

func (stor MemoryStorage) TryGetFullURL(key string) (bool, string) {
	value, found := stor.urls[key]
	return found, value
}

func (stor MemoryStorage) TryGetShortKey(fullURL string) (bool, string) {
	for key, value := range stor.urls {
		if value == fullURL {
			return true, key
		}
	}
	return false, ""
}

func (stor MemoryStorage) SaveURLByShortKey(key string, value string) {
	stor.urls[key] = value
}

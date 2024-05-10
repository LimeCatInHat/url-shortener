package storage

type IURLStogare interface {
	TryGetFullURL(key string) (isSucceed bool, value string)
	TryGetShortKey(fullURL string) (isSucceed bool, value string)
	SaveURLByShortKey(key string, value string)
}

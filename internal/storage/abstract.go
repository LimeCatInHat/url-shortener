package storage

type IUrlStogare interface {
	TryGetFullUrl(key string) (isSucceed bool, value string)
	TryGetShortKey(fullUrl string) (isSucceed bool, value string)
	SaveUrlByShortKey(key string, value string)
}

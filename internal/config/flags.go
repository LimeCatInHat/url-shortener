package config

import (
	"errors"
	"flag"
	"net/url"
	"strconv"
	"strings"
)

type AppConfiguration struct {
	SrvAddr             string
	ShortenLinksBaseURL string
}

var AppSettings = AppConfiguration{
	SrvAddr:             ":8080",
	ShortenLinksBaseURL: "http://localhost:8080/",
}

func SetConfiguration() {
	flag.Func("a", "http server address", func(flagValue string) error {
		value, err := getAddr(flagValue, false)
		if err != nil {
			return err
		}
		AppSettings.SrvAddr = value
		return nil
	})
	flag.Func("b", "shorten url base address", func(flagValue string) error {
		value, err := getAddr(flagValue, true)
		if err != nil {
			return err
		}
		AppSettings.ShortenLinksBaseURL = value
		return nil
	})
	flag.Parse()
}

func getAddr(value string, addTrailingSlashToFullUrl bool) (string, error) {
	portString, isPort := strings.CutPrefix(value, ":")
	if isPort {
		_, err := strconv.Atoi(portString)
		return value, err
	}
	url, err := url.Parse(value)
	isCorrect := err == nil && url.Host != ""
	if !isCorrect {
		return value, errors.New("invalid url")
	}
	result := url.String()
	if !addTrailingSlashToFullUrl || strings.HasSuffix(result, "/") {
		return result, nil
	}

	return result + "/", nil
}

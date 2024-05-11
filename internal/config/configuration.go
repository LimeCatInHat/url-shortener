package config

import (
	"errors"
	"flag"
	"net"
	"net/url"
	"os"
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
	parseFlags()
	setAddressFromServerVariable("SERVER_ADDRESS", &AppSettings.SrvAddr, false)
	setAddressFromServerVariable("BASE_URL", &AppSettings.ShortenLinksBaseURL, true)
}

func parseFlags() {
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

func setAddressFromServerVariable(variableName string, configValueSource *string, needTrailingSlash bool) {
	if envValue := os.Getenv(variableName); envValue != "" {
		result, err := getAddr(envValue, needTrailingSlash)
		if err == nil {
			*configValueSource = result
		}
	}
}

func getAddr(value string, needTrailingSlash bool) (string, error) {
	result, err := getIPBasedAddress(value)
	if err == nil {
		return result, nil
	}
	return getWellFormedRequestURL(value, needTrailingSlash)
}

func getIPBasedAddress(value string) (string, error) {
	parts := strings.Split(value, ":")
	if len(parts) == 1 {
		if isLocalHost(parts[0]) || net.ParseIP(parts[0]) != nil {
			return value, nil
		}
		return value, errors.New("no ip address")
	}
	if len(parts) == 2 {
		_, err := strconv.Atoi(parts[1])
		if err == nil {
			isValid := parts[0] == "" || isLocalHost(parts[0]) || net.ParseIP(parts[0]) != nil
			if isValid {
				return value, nil
			}
		}
	}
	return value, errors.New("no ip address")
}

func getWellFormedRequestURL(value string, needTrailingSlash bool) (string, error) {
	url, err := url.ParseRequestURI(value)
	isCorrect := err == nil && url.Host != ""
	if !isCorrect {
		return value, errors.New("invalid url")
	}
	result := url.String()
	if !needTrailingSlash || strings.HasSuffix(result, "/") {
		return result, nil
	}

	return result + "/", nil
}

func isLocalHost(value string) bool {
	return value == "localhost"
}

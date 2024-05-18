package config

import (
	"flag"
	"fmt"
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

var appConfiguration = AppConfiguration{
	SrvAddr:             ":8080",
	ShortenLinksBaseURL: "http://localhost:8080/",
}

func Init() AppConfiguration {
	parseFlags(&appConfiguration)
	setAddressFromServerVariable("SERVER_ADDRESS", &appConfiguration.SrvAddr, false)
	setAddressFromServerVariable("BASE_URL", &appConfiguration.ShortenLinksBaseURL, true)
	return appConfiguration
}

func GetConfiguration() AppConfiguration {
	return appConfiguration
}

func parseFlags(config *AppConfiguration) {
	flag.Func("a", "http server address", func(flagValue string) error {
		value, err := getAddr(flagValue, false)
		if err != nil {
			return err
		}
		config.SrvAddr = value
		return nil
	})
	flag.Func("b", "shorten url base address", func(flagValue string) error {
		value, err := getAddr(flagValue, true)
		if err != nil {
			return err
		}
		config.ShortenLinksBaseURL = value
		return nil
	})
	flag.Parse()
}

func setAddressFromServerVariable(variableName string, configValueSource *string, needTrailingSlash bool) {
	envValue, exists := os.LookupEnv(variableName)
	if exists {
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
	const maxValidPartsCount = 2
	parts := strings.Split(value, ":")
	if len(parts) == 1 {
		if isLocalHost(parts[0]) || net.ParseIP(parts[0]) != nil {
			return value, nil
		}
		return value, fmt.Errorf(`ip address %q seems to be malformatted`, value)
	}
	if len(parts) == maxValidPartsCount {
		_, err := strconv.Atoi(parts[1])
		if err == nil {
			isValid := parts[0] == "" || isLocalHost(parts[0]) || net.ParseIP(parts[0]) != nil
			if isValid {
				return value, nil
			}
		}
	}
	return value, fmt.Errorf(`ip address %q seems to be malformatted`, value)
}

func getWellFormedRequestURL(value string, needTrailingSlash bool) (string, error) {
	addr, err := url.ParseRequestURI(value)
	isCorrect := err == nil && addr.Host != ""
	if !isCorrect {
		return value, fmt.Errorf(`%q is not valid URL`, value)
	}
	result := addr.String()
	if !needTrailingSlash || strings.HasSuffix(result, "/") {
		return result, nil
	}

	return result + "/", nil
}

func isLocalHost(value string) bool {
	return value == "localhost"
}

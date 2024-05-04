package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/LimeCatInHat/url-shortener/internal/app"
	"github.com/LimeCatInHat/url-shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type searchURLTestDescriptor struct {
	name    string
	request request
	want    response
}
type request struct {
	method string
	path   string
	body   string
}

type response struct {
	statusCode int
	body       string
	headers    map[string]string
}

func TestURLShorterHandler(t *testing.T) {
	tests := []searchURLTestDescriptor{{
		name: "successful shortlink generation",
		request: request{
			method: http.MethodPost,
			body:   "https://yandex.ru/",
			path:   "/",
		},
		want: response{
			statusCode: http.StatusCreated,
			body:       "http://localhost:8080/1e3271ede129813",
			headers: map[string]string{
				"Content-Type": "text/plain",
			},
		},
	}, {
		name: "bad request because of method type",
		request: request{
			method: http.MethodGet,
			path:   "/",
		},
		want: response{
			statusCode: http.StatusBadRequest,
		},
	}, {
		name: "bad request because of empty body",
		request: request{
			method: http.MethodPost,
			path:   "/",
		},
		want: response{
			statusCode: http.StatusBadRequest,
		},
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := createTestRequest(test.request)

			w := httptest.NewRecorder()
			URLShorterHandler(w, request)

			res := w.Result()

			require.Equal(t, test.want.statusCode, res.StatusCode)

			if res.StatusCode <= 299 {
				defer res.Body.Close()
				resBody, err := io.ReadAll(res.Body)
				require.NoError(t, err)
				assert.Equal(t, len(test.want.body) > 0, len(resBody) > 0)

				if test.want.headers != nil {
					for headerKey, headerValue := range test.want.headers {
						assert.Equal(t, headerValue, res.Header.Get(headerKey))
					}
				}
			}
		})
	}
}

func TestSearchFullURLHandler(t *testing.T) {
	tests := []searchURLTestDescriptor{{
		name: "found link",
		request: request{
			method: http.MethodGet,
			path:   "/1e3271ede129813",
		},
		want: response{
			statusCode: http.StatusTemporaryRedirect,
			headers: map[string]string{
				"Location": "https://yandex.ru/",
			},
		},
	}, {
		name: "bad request because of incorrect http method",
		request: request{
			method: http.MethodPost,
			path:   "/1e3271ede129813",
		},
		want: response{
			statusCode: http.StatusBadRequest,
		},
	}, {
		name: "bad request because of empty key",
		request: request{
			method: http.MethodGet,
			path:   "/",
		},
		want: response{
			statusCode: http.StatusBadRequest,
		},
	}, {
		name: "bad request because of to many segments",
		request: request{
			method: http.MethodGet,
			path:   "/1e3271ede129813/extrakey",
		},
		want: response{
			statusCode: http.StatusBadRequest,
		},
	}}

	configureMemoryStorage(map[string]string{"1e3271ede129813": "https://yandex.ru/"})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			request := httptest.NewRequest(test.request.method, test.request.path, nil)

			w := httptest.NewRecorder()
			SearchFullURLHandler(w, request)

			res := w.Result()

			require.Equal(t, test.want.statusCode, res.StatusCode)

			if res.StatusCode <= 299 {
				defer res.Body.Close()
				resBody, err := io.ReadAll(res.Body)
				require.NoError(t, err)
				assert.Equal(t, len(test.want.body) > 0, len(resBody) > 0)

				if test.want.headers != nil {
					for headerKey, headerValue := range test.want.headers {
						assert.Equal(t, headerValue, res.Header.Get(headerKey))
					}
				}
			}
		})
	}
}

func createTestRequest(testRequestData request) *http.Request {
	bodyReader := strings.NewReader(testRequestData.body)
	return httptest.NewRequest(testRequestData.method, testRequestData.path, bodyReader)
}

func configureMemoryStorage(records map[string]string) {
	stor := storage.AppMemoryStorage
	app.ConfigureStorage(stor)
	for itemKey, itemValue := range records {
		stor.SaveURLByShortKey(itemKey, itemValue)
	}
}

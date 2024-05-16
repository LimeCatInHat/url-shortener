package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LimeCatInHat/url-shortener/internal/app"
	"github.com/LimeCatInHat/url-shortener/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-resty/resty/v2"
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
	srv := configureServer(storage.GetStorage())
	defer srv.Close()

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
		name: "only POST method type is supported",
		request: request{
			method: http.MethodGet,
			path:   "/",
		},
		want: response{
			statusCode: http.StatusMethodNotAllowed,
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
		runTest(t, srv, &test)
	}
}

func TestSearchFullURLHandler(t *testing.T) {
	stor := configureMemoryStorage(map[string]string{"1e3271ede129813": "https://yandex.ru/"})
	srv := configureServer(stor)
	defer srv.Close()

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
		name: "only get requests methods allowed",
		request: request{
			method: http.MethodPost,
			path:   "/1e3271ede129813",
		},
		want: response{
			statusCode: http.StatusMethodNotAllowed,
		},
	}, {
		name: "not allowed with empty key",
		request: request{
			method: http.MethodGet,
			path:   "/",
		},
		want: response{
			statusCode: http.StatusMethodNotAllowed,
		},
	}, {
		name: "not found because of to many segments",
		request: request{
			method: http.MethodGet,
			path:   "/1e3271ede129813/extrakey",
		},
		want: response{
			statusCode: http.StatusNotFound,
		},
	}}

	for _, test := range tests {
		runTest(t, srv, &test)
	}
}

func runTest(t *testing.T, srv *httptest.Server, test *searchURLTestDescriptor) {
	t.Helper()
	t.Run(test.name, func(t *testing.T) {
		req := buildRequest(srv, test.request)

		res, err := req.Send()
		require.NoError(t, err)
		require.Equal(t, test.want.statusCode, res.StatusCode())

		if res.StatusCode() <= 299 {
			resBody := res.Body()
			require.NoError(t, err)
			assert.Equal(t, len(test.want.body) > 0, len(resBody) > 0)

			if test.want.headers != nil {
				headers := res.Header()
				for headerKey, headerValue := range test.want.headers {
					assert.Equal(t, headerValue, headers.Get(headerKey))
				}
			}
		}
	})
}

func configureMemoryStorage(records map[string]string) storage.MemoryStorage {
	stor := storage.GetStorage()
	for itemKey, itemValue := range records {
		stor.SaveURLByShortKey(itemKey, itemValue)
	}
	return stor
}

func configureServer(stor app.URLStogare) *httptest.Server {
	r := chi.NewRouter()
	r.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		URLShorterHandler(writer, request, stor)
	})
	r.Get("/{key}", func(writer http.ResponseWriter, request *http.Request) {
		SearchFullURLHandler(writer, request, stor)
	})
	return httptest.NewServer(r)
}

func buildRequest(srv *httptest.Server, testRequest request) *resty.Request {
	client := resty.New().SetRedirectPolicy(noRedirectCustomPolicy())
	req := client.R()
	req.Method = testRequest.method
	req.URL = srv.URL + testRequest.path
	if testRequest.body != "" {
		req.Body = testRequest.body
	}
	return req
}

func noRedirectCustomPolicy() resty.RedirectPolicy {
	return resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	})
}

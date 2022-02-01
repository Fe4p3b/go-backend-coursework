package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Fe4p3b/go-backend-coursework/internal/app/shortener"
	"github.com/Fe4p3b/go-backend-coursework/internal/storage/memory"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_httpHandler_Get(t *testing.T) {
	type fields struct {
		s      shortener.ShortenerService
		method string
		url    string
	}
	type want struct {
		code     int
		response string
		err      bool
	}

	m := memory.New(map[string]string{
		"asdf": "http://yandex.ru",
	})
	s := shortener.New(m, "http://localhost:8080")

	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "test case #1",
			fields: fields{
				s:      s,
				method: http.MethodGet,
				url:    "asdf",
			},
			want: want{
				code:     http.StatusTemporaryRedirect,
				response: "",
				err:      false,
			},
		},
		{
			name: "test case #2",
			fields: fields{
				s:      s,
				method: http.MethodGet,
				url:    "qwerty",
			},
			want: want{
				code:     http.StatusNotFound,
				response: "Not Found",
				err:      true,
			},
		},
		{
			name: "test case #3",
			fields: fields{
				s:      s,
				method: http.MethodGet,
				url:    "",
			},
			want: want{
				code:     http.StatusNotFound,
				response: "Not Found",
				err:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.fields.s)
			request := httptest.NewRequest(tt.fields.method, "/", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(request, w)

			c.SetPath("/:url")
			c.SetParamNames("url")
			c.SetParamValues(tt.fields.url)

			err := h.Get(c)

			if err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.want.code, err.(*echo.HTTPError).Code)
				assert.Equal(t, tt.want.response, err.(*echo.HTTPError).Message)
				return
			}
			// r := chi.NewRouter()
			// r.Get("/{url}", h.GetURL)
			// r.ServeHTTP(w, request)

			assert.NoError(t, err)
			assert.Equal(t, tt.want.code, w.Code)
			assert.Equal(t, tt.want.response, w.Body.String())

		})
	}
}

func Test_httpHandler_post(t *testing.T) {
	type fields struct {
		s      shortener.ShortenerService
		method string
		url    string
		body   string
	}
	type want struct {
		code     int
		response string
		err      bool
	}

	m := memory.New(map[string]string{
		"asdf": "yandex.ru",
	})
	s := shortener.New(m, "http://localhost:8080")

	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "test case #1",
			fields: fields{
				s:      s,
				method: http.MethodPost,
				url:    "/",
				body:   `{"url":"https://yandex.ru"}`,
			},
			want: want{
				code:     http.StatusCreated,
				response: "",
				err:      false,
			},
		},
		{
			name: "test case #2",
			fields: fields{
				s:      s,
				method: http.MethodPost,
				url:    "/",
				body:   `{"url":"yandex.ru"}`,
			},
			want: want{
				code:     http.StatusCreated,
				response: "",
				err:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.fields.s)

			request := httptest.NewRequest(tt.fields.method, tt.fields.url, strings.NewReader(tt.fields.body))
			request.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(request, w)

			err := h.Post(c)
			if tt.want.err {
				assert.Error(t, err)
				assert.Equal(t, tt.want.code, err.(*echo.HTTPError).Code)
				assert.Equal(t, tt.want.response, err.(*echo.HTTPError).Message)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want.code, w.Code)
			if tt.want.response != "" {
				assert.Equal(t, tt.want.response, w.Body.String())
			}
		})
	}
}

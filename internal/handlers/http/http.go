package http

import (
	"errors"
	"net/http"

	"github.com/Fe4p3b/go-backend-coursework/internal/app/shortener"
	"github.com/Fe4p3b/go-backend-coursework/internal/models"
	"github.com/Fe4p3b/go-backend-coursework/internal/storage/memory"
	"github.com/labstack/echo/v4"
)

type handler struct {
	s shortener.ShortenerService
}

func NewHandler(s shortener.ShortenerService) *handler {
	h := &handler{
		s: s,
	}

	return h
}

func (h *handler) Post(c echo.Context) error {
	url := new(models.URL)
	if err := c.Bind(url); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	shortURL, err := h.s.Store(url.URL)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusCreated, &models.URL{ShortURL: shortURL})
}

func (h *handler) Get(c echo.Context) error {
	shortURL := c.Param("url")

	url, err := h.s.Find(shortURL)
	if err != nil {
		if errors.Is(err, memory.ErrorLinkNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		}
		return echo.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

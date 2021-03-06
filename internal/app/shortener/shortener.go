package shortener

import (
	"fmt"

	"github.com/Fe4p3b/go-backend-coursework/internal/repositories"
	"github.com/lithammer/shortuuid"
)

type Service interface {
	Find(string) (string, error)
	Store(string) (string, error)
	AddVisitorCount(string) error
	GetVisitorCounter(string) (int, error)
}

type shortener struct {
	r       repositories.ShortenerRepository
	BaseURL string
}

func New(r repositories.ShortenerRepository, url string) *shortener {
	return &shortener{
		r:       r,
		BaseURL: url,
	}
}

func (s *shortener) Find(url string) (string, error) {
	return s.r.Find(url)
}

func (s *shortener) AddVisitorCount(shorURL string) error {
	return s.r.AddCount(shorURL)
}

func (s *shortener) GetVisitorCounter(shorURL string) (int, error) {
	return s.r.GetVisitorCounter(shorURL)
}

func (s *shortener) Store(url string) (string, error) {
	uuid := shortuuid.New()
	err := s.r.Save(uuid, url)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", s.BaseURL, uuid), nil
}

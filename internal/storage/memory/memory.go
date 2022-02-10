package memory

import (
	"errors"
	"sync"

	"github.com/Fe4p3b/go-backend-coursework/internal/repositories"
)

var ErrorLinkNotFound = errors.New("no such link")
var ErrorDuplicateShortlink = errors.New("duplicate link")
var ErrorMethodIsNotImplemented = errors.New("method is not implemented")
var _ repositories.ShortenerRepository = &memory{}

type memory struct {
	S map[string]string
	C map[string]int
	sync.RWMutex
}

func New(s map[string]string, c map[string]int) *memory {
	return &memory{
		S: s,
		C: c,
	}
}

func (m *memory) Find(url string) (s string, err error) {
	v, ok := m.S[url]
	if !ok {
		return "", ErrorLinkNotFound
	}
	return v, nil
}

func (m *memory) Save(uuid string, url string) error {
	if _, ok := m.S[uuid]; ok {
		return ErrorDuplicateShortlink
	}

	m.Lock()
	m.S[uuid] = url
	m.Unlock()
	return nil
}

func (m *memory) AddCount(shortURL string) error {
	m.Lock()
	m.C[shortURL]++
	m.Unlock()
	return nil
}

func (m *memory) GetVisitorCounter(shortURL string) (int, error) {
	c, ok := m.C[shortURL]
	if !ok {
		return -1, ErrorLinkNotFound
	}

	return c, nil
}

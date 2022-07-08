package entity

import (
	"sync"
)

type DB struct {
	sync.Mutex
	URLs []Link
}

func NewDB() *DB {
	return &DB{}
}

type Link struct {
	shortURL  string
	originURL string
}

const (
	ErrNotFound  = LinkErr("couldn't find the link you were looking for")
	ErrLinkExist = LinkErr("couldn't add link because it already exists")
)

type LinkErr string

func (l LinkErr) Error() string {
	return string(l)
}

type Short map[string]string

func (s Short) SearchLink(shortURL string) (string, error) {
	origin, ok := s[shortURL]
	if !ok {
		return "", ErrNotFound
	}
	return origin, nil
}
func (s Short) AddLink(shortURL, originURL string) error {
	_, err := s.SearchLink(shortURL)
	switch err {
	case ErrNotFound:
		s[shortURL] = originURL
	case nil:
		return ErrLinkExist
	default:
		return err
	}
	return nil
}

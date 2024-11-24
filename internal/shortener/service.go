package shortener

import (
	"crypto/rand"
	"encoding/base64"
)

type Service struct {
	repo URLRepository
}

func NewURLService(repo URLRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Shorten(originalURL string) (string, error) {
	shortURL, err := generateShortURL()
	if err != nil {
		return "", err
	}

	err = s.repo.Save(originalURL, shortURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func (s *Service) GetOriginalURL(shortURL string) (string, error) {
	originalURL, err := s.repo.Get(shortURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}

func generateShortURL() (string, error) {
	b := make([]byte, 6) // 6 bytes for a 8 character base64 string
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

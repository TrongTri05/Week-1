package service

import (
	"ShortLink/repository"
	"math/rand"
	"time"
)

type LinkService struct {
	repo *repository.LinkRepository
}

func NewLinkService(repo *repository.LinkRepository) *LinkService {
	return &LinkService{repo: repo}
}

func GenerateShortCode(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, length)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}

func (s *LinkService) CreateShortLink(originalURL string) (string, error) {
	shortCode := GenerateShortCode(6)
	err := s.repo.Create(shortCode, originalURL)
	return shortCode, err
}

func (s *LinkService) GetOriginalURL(shortCode string) (string, error) {
	url, err := s.repo.FindByShortCode(shortCode)
	if err == nil {
		s.repo.IncreaseClick(shortCode)
	}
	return url, err
}

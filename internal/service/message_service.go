package service

import (
	"consumer-golang/internal/domain"
	"consumer-golang/internal/repository"
)

type MessageService struct {
	repo *repository.MongoRepository
}

func NewMessageService(repo *repository.MongoRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) ProcessMessage(message *domain.Message) error {
	return s.repo.SaveMessage(message)
}

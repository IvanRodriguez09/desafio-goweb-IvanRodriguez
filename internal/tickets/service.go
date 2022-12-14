package tickets

import (
	"context"

	"github.com/IvanRodriguez09/desafio-goweb-IvanRodriguez/internal/domain"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return []domain.Ticket{}, err
	}
	return tickets, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	ticketsDest, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	ticketsAll, err := s.repository.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	return float64(len(ticketsDest)) / float64(len(ticketsAll)) * 100, nil
}

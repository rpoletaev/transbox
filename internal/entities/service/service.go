package service

import (
	"context"
	"transbox/internal/entities"
)

type EntityRepository interface {
	Create(context.Context, entities.Entity) error
}

func New(et EntityRepository) Service {
	return Service{
		entities: et,
	}
}

type Service struct {
	entities EntityRepository
}

func (s Service) Create(ctx context.Context, name string) (entities.Entity, error) {
	et := entities.Entity{Name: name}
	return et, s.entities.Create(ctx, et)
}

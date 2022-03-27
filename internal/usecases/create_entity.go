package usecases

import (
	"context"
	entityRepository "transbox/internal/entities/repository"
	entitiesService "transbox/internal/entities/service"
	outboxProducer "transbox/internal/outbox/producer"
	outboxRepository "transbox/internal/outbox/repository"
	outboxService "transbox/internal/outbox/service"
	"transbox/pkg/eventbus"
	"transbox/pkg/store"
)

func NewCreateEntity(s store.Store) *CreateEntity {
	return &CreateEntity{store: s}
}

type CreateEntity struct {
	store    store.Store
	eventBus eventbus.Bus
}

func (uc *CreateEntity) Create(ctx context.Context, name string) error {
	txWrapper, err := uc.store.Tx(ctx)
	if err != nil {
		return err
	}

	if err := txWrapper.Exec(ctx, func(ctx context.Context, tx store.Querier) error {

		entityRepo := entityRepository.New(tx)
		entitySvc := entitiesService.New(entityRepo)
		createdEntity, err := entitySvc.Create(ctx, name)
		if err != nil {
			return err
		}

		outboxRepo := outboxRepository.New(tx)
		outboxProducer := outboxProducer.New(uc.eventBus)
		outboxSvc := outboxService.New(outboxRepo, outboxProducer)
		if err := outboxSvc.Create(ctx, "EntityCreated", createdEntity.ID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

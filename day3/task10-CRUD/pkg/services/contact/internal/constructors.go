package internal

import (
    "go/pkg/services/contact/internal/delivery"
    "go/pkg/services/contact/internal/repository"
    "go/pkg/services/contact/internal/usecase"
)

func NewContactRepository() repository.ContactRepository {
    return repository.NewContactRepository()
}

func NewGroupRepository() repository.GroupRepository {
    return repository.NewGroupRepository()
}

func NewContactUseCase(contactRepo repository.ContactRepository) usecase.ContactUseCase {
    return usecase.NewContactUseCase(contactRepo)
}

func NewGroupUseCase(groupRepo repository.GroupRepository) usecase.GroupUseCase {
    return usecase.NewGroupUseCase(groupRepo)
}

func NewContactHandler(contactUseCase usecase.ContactUseCase) delivery.ContactHandler {
    return delivery.NewContactHandler(contactUseCase)
}

func NewGroupHandler(groupUseCase usecase.GroupUseCase) delivery.GroupHandler {
    return delivery.NewGroupHandler(groupUseCase)
}

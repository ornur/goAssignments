package usecase

import "go/pkg/services/contact/internal/domain"

type ContactUseCase interface {
    CreateContact(contact *domain.Contact) error
    UpdateContact(contact *domain.Contact) error
    DeleteContact(contactID string) error
    GetContactByID(contactID string) (*domain.Contact, error)
}

type GroupUseCase interface {
    CreateGroup(group *domain.Group) error
    GetGroupByID(groupID string) (*domain.Group, error)
    GetAllGroups() ([]*domain.Group, error)
    AddContactToGroup(contactID, groupID string) error
}

package usecase

import (
    "go/pkg/services/contact/internal/domain"
    "go/pkg/services/contact/internal/repository"
)

type contactUseCaseImpl struct {
    contactRepo repository.ContactRepository
}

func NewContactUseCase(contactRepo repository.ContactRepository) ContactUseCase {
    return &contactUseCaseImpl{
        contactRepo: contactRepo,
    }
}

func (uc *contactUseCaseImpl) CreateContact(contact *domain.Contact) error {
	err := uc.contactRepo.SaveContact(contact)
	if err != nil {
		return err
	}
	return nil
}

func (uc *contactUseCaseImpl) UpdateContact(contact *domain.Contact) error {
	existingContact, err := uc.contactRepo.GetContactByID(contact.ID)
	if err != nil {
		return err
	}

	existingContact.Name = contact.Name
	existingContact.Email = contact.Email
	existingContact.Phone = contact.Phone

	err = uc.contactRepo.SaveContact(existingContact)
	if err != nil {
		return err
	}

	return nil
}

func (uc *contactUseCaseImpl) DeleteContact(contactID string) error {
	err := uc.contactRepo.DeleteContact(contactID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *contactUseCaseImpl) GetContactByID(contactID string) (*domain.Contact, error) {
	contact, err := uc.contactRepo.GetContactByID(contactID)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

type groupUseCaseImpl struct {
    groupRepo repository.GroupRepository
}

func NewGroupUseCase(groupRepo repository.GroupRepository) GroupUseCase {
    return &groupUseCaseImpl{
        groupRepo: groupRepo,
    }
}

func (uc *groupUseCaseImpl) CreateGroup(group *domain.Group) error {
	err := uc.groupRepo.SaveGroup(group)
	if err != nil {
		return err
	}
	return nil
}

func (uc *groupUseCaseImpl) GetGroupByID(groupID string) (*domain.Group, error) {
	group, err := uc.groupRepo.GetGroupByID(groupID)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (uc *groupUseCaseImpl) GetAllGroups() ([]*domain.Group, error) {
	groups, err := uc.groupRepo.GetAllGroups()
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (uc *groupUseCaseImpl) AddContactToGroup(contactID, groupID string) error {
	contact, err := uc.contactRepo.GetContactByID(contactID)
	if err != nil {
		return err
	}

	group, err := uc.groupRepo.GetGroupByID(groupID)
	if err != nil {
		return err
	}

	group.Contacts = append(group.Contacts, contact)

	err = uc.groupRepo.SaveGroup(group)
	if err != nil {
		return err
	}

	return nil
}

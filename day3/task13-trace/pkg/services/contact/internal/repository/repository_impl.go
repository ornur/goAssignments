package repository

import (
	"database/sql"
	"go/pkg/services/contact/internal/domain"
)

type contactRepositoryImpl struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) ContactRepository {
	return &contactRepositoryImpl{
		db: db,
	}
}

func (r *contactRepositoryImpl) SaveContact(contact *domain.Contact) error {
	query := "INSERT INTO contacts (name, email) VALUES (?, ?)"
	_, err := r.db.Exec(query, contact.Name, contact.Email)
	return err
}

func (r *contactRepositoryImpl) GetContactByID(contactID string) (*domain.Contact, error) {
	query := "SELECT name, email FROM contacts WHERE id = ?"
	row := r.db.QueryRow(query, contactID)

	contact := &domain.Contact{}
	err := row.Scan(&contact.Name, &contact.Email)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (r *contactRepositoryImpl) DeleteContact(contactID string) error {
	query := "DELETE FROM contacts WHERE id = ?"
	_, err := r.db.Exec(query, contactID)
	if err != nil {
		return err
	}
	return nil
}

type groupRepositoryImpl struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) GroupRepository {
	return &groupRepositoryImpl{
		db: db,
	}
}

func (r *groupRepositoryImpl) SaveGroup(group *domain.Group) error {
	query := "INSERT INTO groups (name) VALUES (?)"
	_, err := r.db.Exec(query, group.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *groupRepositoryImpl) GetGroupByID(groupID string) (*domain.Group, error) {
	query := "SELECT name FROM groups WHERE id = ?"
	row := r.db.QueryRow(query, groupID)

	group := &domain.Group{}
	err := row.Scan(&group.Name)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (r *groupRepositoryImpl) GetAllGroups() ([]*domain.Group, error) {
	query := "SELECT name FROM groups"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	groups := []*domain.Group{}
	for rows.Next() {
		group := &domain.Group{}
		err := rows.Scan(&group.Name)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *groupRepositoryImpl) AddContactToGroup(contact *domain.Contact, groupID string) error {
	query := "INSERT INTO group_contacts (group_id, contact_id) VALUES (?, ?)"
	_, err := r.db.Exec(query, groupID, contact.ID)
	if err != nil {
		return err
	}
	return nil
}

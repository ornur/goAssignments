package repository

import (
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

// SaveContact saves a contact
func (r *contactRepositoryImpl) SaveContact(contact *domain.Contact) error {
    query := "INSERT INTO contacts (name, email) VALUES (?, ?)"
    _, err := r.db.Exec(query, contact.Name, contact.Email)
    return err
}

// GetContactByID retrieves a contact by ID
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

// DeleteContact deletes a contact by ID
func (r *contactRepositoryImpl) DeleteContact(contactID string) error {
	query := "DELETE FROM contacts WHERE id = ?"
	_, err := r.db.Exec(query, contactID)
	if err != nil {
		return err
	}
	return nil
}

// groupRepositoryImpl represents the repository implementation for group operations
type groupRepositoryImpl struct {
    db *sql.DB
}

// NewGroupRepository creates a new instance of GroupRepository
func NewGroupRepository(db *sql.DB) GroupRepository {
    return &groupRepositoryImpl{
        db: db,
    }
}

// SaveGroup saves a group
func (r *groupRepositoryImpl) SaveGroup(group *domain.Group) error {
	query := "INSERT INTO groups (name) VALUES (?)"
	_, err := r.db.Exec(query, group.Name)
	if err != nil {
		return err
	}
	return nil
}

// GetGroupByID retrieves a group by ID
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

// GetAllGroups retrieves all groups
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

// AddContactToGroup adds a contact to a group
func (r *groupRepositoryImpl) AddContactToGroup(contact *domain.Contact, groupID string) error {
	query := "INSERT INTO group_contacts (group_id, contact_id) VALUES (?, ?)"
	_, err := r.db.Exec(query, groupID, contact.ID)
	if err != nil {
		return err
	}
	return nil
}

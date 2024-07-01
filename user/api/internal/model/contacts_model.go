package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

var _ ContactsModel = (*customContactsModel)(nil)

type (
	// ContactsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContactsModel.
	ContactsModel interface {
		contactsModel
		customContactsLogicModel
	}

	customContactsModel struct {
		*defaultContactsModel
	}

	customContactsLogicModel interface {
		QueryUsers(ctx context.Context, query string) ([]Contacts, error)
		QueryUser(ctx context.Context, query string) (*Contacts, error)
	}
)

func (m *defaultContactsModel) QueryUser(ctx context.Context, query string) (*Contacts, error) {
	var contact *Contacts
	err := m.conn.WithContext(ctx).Where(query).Find(&contact).Error
	switch {
	case err == nil:
		return contact, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultContactsModel) QueryUsers(ctx context.Context, query string) ([]Contacts, error) {
	var contacts []Contacts
	err := m.conn.WithContext(ctx).Where(query).Find(&contacts).Error
	switch {
	case err == nil:
		return contacts, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewContactsModel returns a model for the database table.
func NewContactsModel(conn *gorm.DB) ContactsModel {
	return &customContactsModel{
		defaultContactsModel: newContactsModel(conn),
	}
}

func (m *defaultContactsModel) customCacheKeys(data *Contacts) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

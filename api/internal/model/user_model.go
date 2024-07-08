package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		customUserLogicModel
	}

	customUserModel struct {
		*defaultUserModel
	}

	customUserLogicModel interface {
		QueryUser(ctx context.Context, query string) (*User, error)
		QueryUsersByKey(ctx context.Context, query string, uid int) ([]User, error)
	}
)

func (m *defaultUserModel) QueryUsersByKey(ctx context.Context, query string, uid int) ([]User, error) {
	var user []User
	// select * from user where id in
	//(select contact_id
	//from contacts where owner_id=2) and nickname like '%仁%';
	err := m.conn.WithContext(ctx).Where(query).
		Where("id IN (?)",
			m.conn.Table("contacts").Select("contact_id").Where("owner_id = ?", uid),
		).Find(&user).Error
	switch {
	case err == nil:
		return user, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) QueryUser(ctx context.Context, query string) (*User, error) {
	var user User
	err := m.conn.WithContext(ctx).Where(query).Find(&user).Error
	switch {
	case err == nil:
		return &user, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewUserModel returns a model for the database table.
func NewUserModel(conn *gorm.DB) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *defaultUserModel) customCacheKeys(data *User) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

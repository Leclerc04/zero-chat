package model

import (
	"context"
	"gorm.io/gorm"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel
		customMessageLogicModel
	}

	customMessageModel struct {
		*defaultMessageModel
	}

	customMessageLogicModel interface {
		GetAllHistory(ctx context.Context)
	}
)

func (m *defaultMessageModel) GetAllHistory(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

// NewMessageModel returns a model for the database table.
func NewMessageModel(conn *gorm.DB) MessageModel {
	return &customMessageModel{
		defaultMessageModel: newMessageModel(conn),
	}
}

func (m *defaultMessageModel) customCacheKeys(data *Message) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

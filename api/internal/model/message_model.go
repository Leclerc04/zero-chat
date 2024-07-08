package model

import (
	"context"
	"gorm.io/gorm"
	"log"
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
		GetAllHistory(ctx context.Context, uid int64) ([]Message, error)
		GetDetailHistory(ctx context.Context, toUid, fromUid int64) ([]Message, error)
	}
)

func (m *defaultMessageModel) GetDetailHistory(ctx context.Context, toUid, fromUid int64) ([]Message, error) {
	var msgs []Message
	if err := m.conn.WithContext(ctx).
		Raw("select * from message where (send_id = ? and receive_id = ?) or (send_id = ? and receive_id = ?) order by created_at", toUid, fromUid, fromUid, toUid).
		Scan(&msgs).Error; err != nil {
		return nil, err
	}
	return msgs, nil
}

func (m *defaultMessageModel) GetAllHistory(ctx context.Context, uid int64) ([]Message, error) {
	var msgs []Message
	var res []int64
	if err := m.conn.WithContext(ctx).Raw(
		"SELECT DISTINCT contact_id FROM "+
			"(SELECT send_id AS contact_id FROM message WHERE receive_id = ? UNION SELECT receive_id AS contact_id FROM message WHERE send_id = ?) AS contacts "+
			"WHERE contact_id != ?", uid, uid, uid).Scan(&res).Error; err != nil {
		log.Printf("err:%s", err)
		return nil, err
	}
	for _, re := range res {
		tmp := Message{
			Id:  re,
			Msg: m.queryLastMSg(ctx, re),
		}
		msgs = append(msgs, tmp)
	}
	return msgs, nil
}
func (m *defaultMessageModel) queryLastMSg(ctx context.Context, uid int64) string {
	var message Message
	if err := m.conn.WithContext(ctx).Where("send_id=? or receive_id=?", uid, uid).Order("created_at desc").Find(&message).Error; err != nil {
		log.Printf("error:%s", err)
		return ""
	}
	return message.Msg
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

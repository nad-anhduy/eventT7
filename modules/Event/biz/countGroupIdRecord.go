package biz

import "eventT7/modules/Event/model"

type CountGroupIdRecordStorage interface {
	GetAmountByGroupId() (*[]model.AmountGroupIdStructToDB, error)
}

type countGroupIdRecordBiz struct {
	store CountGroupIdRecordStorage
}

func NewCountGroupIdRecord(store CountGroupIdRecordStorage) *countGroupIdRecordBiz {
	return &countGroupIdRecordBiz{store: store}
}

func (c *countGroupIdRecordBiz) CountNewGroupIdRecord() (*[]model.AmountGroupIdStructToDB, error) {
	result, err := c.store.GetAmountByGroupId()
	if err != nil {
		return nil, err
	}
	return result, nil
}

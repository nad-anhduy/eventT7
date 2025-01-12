package biz

import (
	"context"
	"eventT7/modules/Event/model"
)

type CreateRecordStorage interface {
	GetRecordExistBySession(ctx context.Context, session, userName, groupID string) (bool, error)
	InsertData(ctx context.Context, data *model.InsertStructToDB) error
}

type createRecordBiz struct {
	store CreateRecordStorage
}

func NewCreateRecord(store CreateRecordStorage) *createRecordBiz {
	return &createRecordBiz{store: store}
}

func (c *createRecordBiz) CreateNewRecord(ctx context.Context, data model.InsertStructToDB) error {
	ok, err := c.store.GetRecordExistBySession(ctx, data.Session, data.UserName, data.GroupID)
	if ok && err.Error() == "record already exists" {
		return err
	}

	err = c.store.InsertData(ctx, &data)
	if err != nil {
		return err
	}

	return nil
}

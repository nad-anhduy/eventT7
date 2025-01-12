package biz

import (
	"context"
	"errors"
	"eventT7/modules/Event/model"
)

type UpdateRecordStorage interface {
	GetRecordBySession(ctx context.Context, session string) (*model.ResponseStructFromDB, error)
	UpdateData(c context.Context, uid string, data *model.UpdateStructToDB) error
}

type updateRecordBiz struct {
	store UpdateRecordStorage
}

func NewUpdateRecord(store UpdateRecordStorage) *updateRecordBiz {
	return &updateRecordBiz{store: store}
}

func (u *updateRecordBiz) UpdateNewRecord(ctx context.Context, data model.UpdateStructToDB) error {
	resultCheck, err := u.store.GetRecordBySession(ctx, data.Session)
	if err != nil || resultCheck == nil {
		return errors.New(`Record not found`)
	}

	err = u.store.UpdateData(ctx, resultCheck.Uid, &data)
	if err != nil {
		return err
	}
	return nil
}

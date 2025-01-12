package storage

import (
	"context"
	"errors"
	"eventT7/modules/Event/model"
)

func (s *sqlStore) GetRecordBySession(ctx context.Context, session string) (*model.ResponseStructFromDB, error) {
	var result model.ResponseStructFromDB

	if err := s.db.Table(model.ResponseStructFromDB{}.TableName()).Where(map[string]interface{}{"session": session}).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *sqlStore) GetRecordExistBySession(ctx context.Context, session, userName, groupID string) (bool, error) {
	var result model.ResponseStructFromDB

	if err := s.db.Table(model.ResponseStructFromDB{}.TableName()).Where(map[string]interface{}{"session": session, "userName": userName, "unVote": false}).First(&result).Error; err != nil {
		return false, err
	}
	if result.Uid != "" {
		return true, errors.New("record already exists")
	}
	return true, errors.New("record don't exists")
}

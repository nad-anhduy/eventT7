package storage

import (
	"context"
	"eventT7/modules/Event/model"
)

func (s *sqlStore) InsertData(ctx context.Context, data *model.InsertStructToDB) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

package storage

import (
	"context"
	"eventT7/modules/Event/model"
)

func (s *sqlStore) UpdateData(c context.Context, uid string, data *model.UpdateStructToDB) error {
	if err := s.db.Where(map[string]interface{}{"uid": uid}).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

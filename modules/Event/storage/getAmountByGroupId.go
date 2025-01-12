package storage

import (
	"errors"
	"eventT7/modules/Event/model"
)

func (s *sqlStore) GetAmountByGroupId() (*[]model.AmountGroupIdStructToDB, error) {

	var result []model.AmountGroupIdStructToDB
	if err := s.db.Raw(`select "session", "groupId", "img",COUNT(*) AS "SL" from event_soul_history WHERE "unVote" = FALSE and "session" <> 'session1' GROUP BY "groupId", "session", "img" ORDER BY "SL" DESC LIMIT 200`).Scan(&result).Error; err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New(`data after scan = 0`)
	}

	return &result, nil

}

package storage

import (
	"eventT7/modules/Event/model"
	"fmt"
)

func (s *sqlStore) GetListSession() ([]string, error) {
	var list []string

	if err := s.db.Raw(`select "session" from event_soul_history GROUP BY "session"`).Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *sqlStore) GetTopPerSession(nameSession string) (*[]model.TopVotePerSessionStructToDB, error) {
	var listTop []model.TopVotePerSessionStructToDB

	query := fmt.Sprintf(`select "groupId", "extraX",COUNT(*) AS "SL" from event_soul_history WHERE event_soul_history.session = '%s' and "unVote" = FALSE and "groupId" <> 'session1' GROUP BY "groupId", "extraX" ORDER BY "SL" DESC LIMIT 3`, nameSession)
	if err := s.db.Raw(query).Scan(&listTop).Error; err != nil {
		return nil, err
	}
	return &listTop, nil
}

package biz

import (
	"errors"
	"eventT7/modules/Event/model"
	"log"
	"time"
)

func ProxyFilterByTimeRequest(id, groupId string) error {

	if val, ok := model.ListGroupId[groupId]; ok {
		if time.Now().UnixMilli()-val < 2000 {
			log.Printf(`SoulT7 - [%s] - Id: %s too many request`, id, groupId)
			model.ListGroupId[groupId] = time.Now().UnixMilli()
			return errors.New(`Limit request`)
		}
	} else {
		model.ListGroupId[groupId] = time.Now().UnixMilli()
	}
	return nil
}

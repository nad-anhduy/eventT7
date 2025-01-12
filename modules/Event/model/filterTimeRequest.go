package model

import (
	"errors"
	"log"
	"time"
)

func ProxyFilterByTimeRequest(id, groupId string) error {

	if val, ok := ListGroupId[groupId]; ok {
		if time.Now().UnixMilli()-val < 2000 {
			log.Printf(`SoulT7 - [%s] - Id: %s too many request`, id, groupId)
			ListGroupId[groupId] = time.Now().UnixMilli()
			return errors.New(`Limit request`)
		}
	} else {
		ListGroupId[groupId] = time.Now().UnixMilli()
	}
	return nil
}

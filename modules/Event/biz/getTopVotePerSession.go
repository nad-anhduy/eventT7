package biz

import "eventT7/modules/Event/model"

type TopVotePerSessionStorage interface {
	GetListSession() ([]string, error)
	GetTopPerSession(nameSession string) (*[]model.TopVotePerSessionStructToDB, error)
}

type topVotePerSessionBiz struct {
	store TopVotePerSessionStorage
}

func NewTopVotePerSession(store TopVotePerSessionStorage) *topVotePerSessionBiz {
	return &topVotePerSessionBiz{store: store}
}

func (biz *topVotePerSessionBiz) NewListSession() ([]string, error) {

	listSession, err := biz.store.GetListSession()
	if err != nil {
		return nil, err
	}
	return listSession, nil
}

func (biz *topVotePerSessionBiz) TopVoteNewSession(name string) (*[]model.TopVotePerSessionStructToDB, error) {

	result, err := biz.store.GetTopPerSession(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

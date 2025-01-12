package model

import "eventT7/common"

type RequestStruct struct {
	UserName  UserName `json:"user_name"`
	GroupID   GroupID  `json:"group_id"`
	Session   Session  `json:"session"`
	ImgUrl    ImgUrl   `json:"img_url"`
	ExtraData []string `json:"extra_data"`
}

type ImgUrl string

func (u ImgUrl) String() string { return string(u) }

func (u ImgUrl) Valid() error {
	err := common.Validate(u.String())
	if err != nil {
		return err
	}
	return nil
}

type UserName string

func (u UserName) String() string { return string(u) }

func (u UserName) Valid() error {
	err := common.Validate(u.String())
	if err != nil {
		return err
	}
	return nil
}

type GroupID string

func (u GroupID) String() string { return string(u) }

func (u GroupID) Valid() error {
	err := common.Validate(u.String())
	if err != nil {
		return err
	}
	return nil
}

type Session string

func (u Session) String() string { return string(u) }

func (u Session) Valid() error {
	err := common.Validate(u.String())
	if err != nil {
		return err
	}
	return nil
}

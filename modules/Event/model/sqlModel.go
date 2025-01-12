package model

var ListGroupDeny []string
var ListGroupId = make(map[string]int64)

const tableName = "event_soul_history"

type ResponseStructFromDB struct {
	ID       int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Created  int64  `gorm:"column:created"`
	Updated  int64  `gorm:"column:updated"`
	Uid      string `gorm:"column:uid"`
	Session  string `gorm:"column:session"`
	UserName string `gorm:"column:userName"`
	GroupID  string `gorm:"column:groupId"`
	ExtraX   string `gorm:"column:extraX"`
}

func (ResponseStructFromDB) TableName() string { return tableName }

type InsertStructToDB struct {
	ID       int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Created  int64  `gorm:"column:created"`
	Uid      string `gorm:"column:uid"`
	Session  string `gorm:"column:session"`
	UserName string `gorm:"column:userName"`
	GroupID  string `gorm:"column:groupId"`
	ImgUrl   string `gorm:"column:img"`
	ExtraX   string `gorm:"column:extraX"`
}

func (InsertStructToDB) TableName() string { return tableName }

type UpdateStructToDB struct {
	ID      int   `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Updated int64 `gorm:"column:updated"`
	Session string
	UnVote  bool `gorm:"column:unVote"`
}

func (UpdateStructToDB) TableName() string { return tableName }

type AmountGroupIdStructToDB struct {
	Session string `gorm:"column:session" json:"session"`
	GroupID string `gorm:"column:groupId" json:"group_id"`
	ImgUrl  string `gorm:"column:img" json:"img_url"`
	Amount  int64  `gorm:"column:SL" json:"amount"`
}

type TopVotePerSessionStructToDB struct {
	ExtraX  string `gorm:"column:extraX" json:"extra_data"`
	GroupID string `gorm:"column:groupId" json:"group_id"`
	Amount  int64  `gorm:"column:SL" json:"amount"`
}

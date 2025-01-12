package model

type ResponseStruct struct {
	Uid       string `json:"id"`
	UserName  string `json:"user_name"`
	GroupID   string `json:"group_id"`
	Session   string `json:"session"`
	ExtraData string `json:"extra_data"`
	Created   int64  `json:"created"`
	Updated   int64  `json:"updated"`
}

type ResponseTopVoteStruct struct {
	Session string                         `json:"session"`
	Items   *[]TopVotePerSessionStructToDB `json:"items"`
}

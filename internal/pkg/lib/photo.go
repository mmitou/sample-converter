package lib

type ReactionCounter struct {
	Plus1      int `json:"+1"`
	Smile      int `json:"smile"`
	Sunglasses int `json:"sunglasses"`
	Heart      int `json:"heart"`
}

type Photo struct {
	PK        string
	SK        string
	Username  string          `json:"username"`
	Timestamp string          `json:"timestamp"`
	Location  string          `json:"location"`
	Reactions ReactionCounter `json:"reactions"`
}

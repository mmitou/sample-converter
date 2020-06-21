package lib

type Reaction struct {
	PK           string
	SK           string
	ReactingUser string `json:"reactingUser"`
	Photo        string `json:"photo"`
	ReactionType string `json:"reactionType"`
	Timestamp    string `json:"timestamp"`
}

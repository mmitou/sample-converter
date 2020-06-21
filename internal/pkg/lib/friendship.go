package lib

type Friendship struct {
	PK            string
	SK            string
	FollowedUser  string `json:"followdUser"`
	FollowingUser string `json:"followingUser"`
	Timestamp     string `json:"timestamp"`
}

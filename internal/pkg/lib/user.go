package lib

type User struct {
	PK          string
	SK          string
	Address     string   `json:"address"`
	Birthdate   string   `json:"birthdate"`
	Email       string   `json:"email"`
	Name        string   `json:"name"`
	Username    string   `json:"username"`
	Status      string   `json:"status"`
	Interests   []string `json:"interests"`
	Followers   int      `json:"followers"`
	Following   int      `json:"following"`
	PinnedImage string   `json:"pinnedImage"`
}

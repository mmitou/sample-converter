package lib

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
)

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

func NewUserFacet(users []*User) *Facet {
	var attrs []map[string]*dynamodb.AttributeValue
	for _, user := range users {
		attr, err := dynamo.MarshalItem(user)
		if err != nil {
			panic(err)
		}
		attrs = append(attrs, attr)
	}

	facet := &Facet{
		FacetName:         "User",
		KeyAttributeAlias: KeyAlias{PartitionKeyAlias: "PK", SortKeyAlias: "SK"},
		TableData:         attrs,
		NonKeyAttributes:  NonKeyAttributes(User{}),
		DataAccess:        DataAccessor{map[string]interface{}{}},
	}

	return facet
}

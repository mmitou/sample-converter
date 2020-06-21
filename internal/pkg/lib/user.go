package lib

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
)

type User struct {
	PK          string
	SK          string
	Address     string   `json:"address" dynamo:"address"`
	Birthdate   string   `json:"birthdate" dynamo:"birthdate"`
	Email       string   `json:"email" dynamo:"email"`
	Name        string   `json:"name" dynamo:"name"`
	Username    string   `json:"username" dynamo:"username"`
	Status      string   `json:"status" dynamo:"status"`
	Interests   []string `json:"interests" dynamo:"interests"`
	Followers   int      `json:"followers" dynamo:"followers"`
	Following   int      `json:"following" dynamo:"following"`
	PinnedImage string   `json:"pinnedImage" dynamo:"pinnedImage"`
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

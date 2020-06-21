package lib

import (
	"reflect"

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

func UserNonKeyAttributes() []string {
	u := User{}
	t := reflect.TypeOf(u)
	var tags []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if len(tag) != 0 {
			tags = append(tags, tag)
		}
	}
	return tags
}

func toAttributeValues(objs []interface{}) []*dynamodb.AttributeValue {
	var attrs []*dynamodb.AttributeValue
	for _, obj := range objs {
		attr, err := dynamo.Marshal(obj)
		if err != nil {
			panic(err)
		}
		attrs = append(attrs, attr)
	}
	return attrs
}

func NewUserFacet(users []*User) *Facet {
	var attrs []*dynamodb.AttributeValue
	for _, user := range users {
		attr, err := dynamo.Marshal(user)
		if err != nil {
			panic(err)
		}
		attrs = append(attrs, attr)
	}

	facet := &Facet{
		FacetName:         "User",
		KeyAttributeAlias: KeyAlias{PartitionKeyAlias: "PK", SortKeyAlias: "SK"},
		TableData:         attrs,
		NonKeyAttributes:  UserNonKeyAttributes(),
		DataAccess:        DataAccessor{map[string]interface{}{}},
	}

	return facet
}

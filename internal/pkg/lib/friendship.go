package lib

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
)

type Friendship struct {
	PK            string
	SK            string
	FollowedUser  string `json:"followdUser"`
	FollowingUser string `json:"followingUser"`
	Timestamp     string `json:"timestamp"`
}

func NewFriendshipFacet(fs []*Friendship) *Facet {
	var attrs []map[string]*dynamodb.AttributeValue
	for _, f := range fs {
		attr, err := dynamo.MarshalItem(f)
		if err != nil {
			panic(err)
		}
		attrs = append(attrs, attr)
	}

	facet := &Facet{
		FacetName:         "Friendship",
		KeyAttributeAlias: KeyAlias{PartitionKeyAlias: "PK", SortKeyAlias: "SK"},
		TableData:         attrs,
		NonKeyAttributes:  NonKeyAttributes(Friendship{}),
		DataAccess:        DataAccessor{map[string]interface{}{}},
	}

	return facet

}

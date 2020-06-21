package lib

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
)

type Reaction struct {
	PK           string
	SK           string
	ReactingUser string `json:"reactingUser"`
	Photo        string `json:"photo"`
	ReactionType string `json:"reactionType"`
	Timestamp    string `json:"timestamp"`
}

func NewReactionFacet(rs []*Reaction) *Facet {
	var attrs []map[string]*dynamodb.AttributeValue
	for _, r := range rs {
		attr, err := dynamo.MarshalItem(r)
		if err != nil {
			panic(err)
		}
		attrs = append(attrs, attr)
	}

	facet := &Facet{
		FacetName:         "Photo",
		KeyAttributeAlias: KeyAlias{PartitionKeyAlias: "PK", SortKeyAlias: "SK"},
		TableData:         attrs,
		NonKeyAttributes:  NonKeyAttributes(Photo{}),
		DataAccess:        DataAccessor{map[string]interface{}{}},
	}

	return facet
}

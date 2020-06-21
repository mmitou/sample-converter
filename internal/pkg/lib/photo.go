package lib

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
)

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

func NewPhotoFacet(ps []*Photo) *Facet {
	var attrs []map[string]*dynamodb.AttributeValue
	for _, p := range ps {
		attr, err := dynamo.MarshalItem(p)
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

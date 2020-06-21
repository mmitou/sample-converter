package lib

import (
	"reflect"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type KeyAlias struct {
	PartitionKeyAlias string
	SortKeyAlias      string
}

type DataAccessor struct {
	MySql map[string]interface{}
}

type Facet struct {
	FacetName         string
	KeyAttributeAlias KeyAlias
	TableData         []map[string]*dynamodb.AttributeValue
	NonKeyAttributes  []string
	DataAccess        DataAccessor
}

func NonKeyAttributes(obj interface{}) []string {
	t := reflect.TypeOf(obj)
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

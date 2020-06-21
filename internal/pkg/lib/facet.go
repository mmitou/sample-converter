package lib

import "github.com/aws/aws-sdk-go/service/dynamodb"

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
	TableData         []*dynamodb.AttributeValue
	NonKeyAttributes  []string
	DataAccess        DataAccessor
}

package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ModelMetadata struct {
	Author           string
	DateCreated      string
	DateLastMofified string
	Description      string
	Version          string
}

type KeyAttributes struct {
	PartitionKey *dynamodb.AttributeDefinition
	SortKey      *dynamodb.AttributeDefinition
}

func NewKeyAttributes() *KeyAttributes {
	k := &KeyAttributes{
		PartitionKey: &dynamodb.AttributeDefinition{
			AttributeName: aws.String("PK"),
			AttributeType: aws.String("S"),
		},
		SortKey: &dynamodb.AttributeDefinition{
			AttributeName: aws.String("SK"),
			AttributeType: aws.String("S"),
		},
	}
	return k
}

type Table struct {
	TableName        string
	KeyAttributes    *KeyAttributes
	NonKeyAttributes []*dynamodb.AttributeDefinition
	TableFacets      []*Facet
}

type DataModel struct {
	ModelName     string
	ModelMetadata ModelMetadata
	DataModel     []*Table
}

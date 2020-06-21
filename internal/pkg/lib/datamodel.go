package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ModelMetadata struct {
	Author           string
	DateCreated      string
	DateLastModified string
	Description      string
	Version          string
}

func NewModelMetadata() *ModelMetadata {
	data := &ModelMetadata{
		Author:           "mmitou",
		DateCreated:      "Jun 21, 2020, 12:09 PM",
		DateLastModified: "Jun 21, 2020, 12:14 PM",
		Description:      "test",
		Version:          "1.0",
	}
	return data
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

func NewNonKeyAttributes() []*dynamodb.AttributeDefinition {
	xs := []*dynamodb.AttributeDefinition{
		{AttributeName: aws.String("address"), AttributeType: aws.String("S")},
		{AttributeName: aws.String("birthdate"), AttributeType: aws.String("S")},
		{AttributeName: aws.String("email"), AttributeType: aws.String("S")},
		{AttributeName: aws.String("name"), AttributeType: aws.String("S")},
		{AttributeName: aws.String("username"), AttributeType: aws.String("S")},
		{AttributeName: aws.String("status"), AttributeType: aws.String("S")},
		{AttributeName: aws.String("interests"), AttributeType: aws.String("L")},
		{AttributeName: aws.String("followers"), AttributeType: aws.String("N")},
		{AttributeName: aws.String("following"), AttributeType: aws.String("N")},
		{AttributeName: aws.String("pinnedImage"), AttributeType: aws.String("S")},
		{AttributeName: aws.String("timestamp"), AttributeType: aws.String("S")},
		{AttributeName: aws.String("location"), AttributeType: aws.String("S")},
		{AttributeName: aws.String("reactions"), AttributeType: aws.String("M")},
		{AttributeName: aws.String("followdUser"), AttributeType: aws.String("S")},
		{AttributeName: aws.String("followingUser"), AttributeType: aws.String("S")},
	}
	return xs
}

type Table struct {
	TableName        string
	KeyAttributes    *KeyAttributes
	NonKeyAttributes []*dynamodb.AttributeDefinition
	TableFacets      []*Facet
}

func NewTable() *Table {
	tbl := &Table{
		TableName:        "quick-photos",
		KeyAttributes:    NewKeyAttributes(),
		NonKeyAttributes: NewNonKeyAttributes(),
	}
	return tbl
}

type DataModel struct {
	ModelName     string
	ModelMetadata *ModelMetadata
	DataModel     []*Table
}

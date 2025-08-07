package config

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DynamoDB table names
const (
	UsersTableName           = "LiloUsers"
	StyleProfilesTableName   = "LiloStyleProfiles"
	ClothingItemsTableName   = "LiloClothingItems"
	OutfitsTableName         = "LiloOutfits"
	ReflectionsTableName     = "LiloReflections"
	RecommendationsTableName = "LiloRecommendations"
)

// CreateDynamoDBTables creates all required DynamoDB tables if they don't exist
func CreateDynamoDBTables(client *dynamodb.Client) error {
	tables := []struct {
		Name         string
		KeySchema    []types.KeySchemaElement
		AttributeDef []types.AttributeDefinition
		GSIs         []types.GlobalSecondaryIndex
	}{
		{
			Name: UsersTableName,
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:       types.KeyTypeHash,
				},
			},
			AttributeDef: []types.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("email"),
					AttributeType: types.ScalarAttributeTypeS,
				},
			},
			GSIs: []types.GlobalSecondaryIndex{
				{
					IndexName: aws.String("EmailIndex"),
					KeySchema: []types.KeySchemaElement{
						{
							AttributeName: aws.String("email"),
							KeyType:       types.KeyTypeHash,
						},
					},
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
					ProvisionedThroughput: &types.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(5),
						WriteCapacityUnits: aws.Int64(5),
					},
				},
			},
		},
		{
			Name: StyleProfilesTableName,
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:       types.KeyTypeHash,
				},
			},
			AttributeDef: []types.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("userId"),
					AttributeType: types.ScalarAttributeTypeS,
				},
			},
			GSIs: []types.GlobalSecondaryIndex{
				{
					IndexName: aws.String("UserIdIndex"),
					KeySchema: []types.KeySchemaElement{
						{
							AttributeName: aws.String("userId"),
							KeyType:       types.KeyTypeHash,
						},
					},
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
					ProvisionedThroughput: &types.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(5),
						WriteCapacityUnits: aws.Int64(5),
					},
				},
			},
		},
		{
			Name: ClothingItemsTableName,
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:       types.KeyTypeHash,
				},
			},
			AttributeDef: []types.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("userId"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("category"),
					AttributeType: types.ScalarAttributeTypeS,
				},
			},
			GSIs: []types.GlobalSecondaryIndex{
				{
					IndexName: aws.String("UserIdIndex"),
					KeySchema: []types.KeySchemaElement{
						{
							AttributeName: aws.String("userId"),
							KeyType:       types.KeyTypeHash,
						},
					},
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
					ProvisionedThroughput: &types.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(5),
						WriteCapacityUnits: aws.Int64(5),
					},
				},
				{
					IndexName: aws.String("UserCategoryIndex"),
					KeySchema: []types.KeySchemaElement{
						{
							AttributeName: aws.String("userId"),
							KeyType:       types.KeyTypeHash,
						},
						{
							AttributeName: aws.String("category"),
							KeyType:       types.KeyTypeRange,
						},
					},
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
					ProvisionedThroughput: &types.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(5),
						WriteCapacityUnits: aws.Int64(5),
					},
				},
			},
		},
		{
			Name: OutfitsTableName,
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:       types.KeyTypeHash,
				},
			},
			AttributeDef: []types.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("userId"),
					AttributeType: types.ScalarAttributeTypeS,
				},
			},
			GSIs: []types.GlobalSecondaryIndex{
				{
					IndexName: aws.String("UserIdIndex"),
					KeySchema: []types.KeySchemaElement{
						{
							AttributeName: aws.String("userId"),
							KeyType:       types.KeyTypeHash,
						},
					},
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
					ProvisionedThroughput: &types.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(5),
						WriteCapacityUnits: aws.Int64(5),
					},
				},
			},
		},
		{
			Name: ReflectionsTableName,
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:       types.KeyTypeHash,
				},
			},
			AttributeDef: []types.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("userId"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("outfitId"),
					AttributeType: types.ScalarAttributeTypeS,
				},
			},
			GSIs: []types.GlobalSecondaryIndex{
				{
					IndexName: aws.String("UserIdIndex"),
					KeySchema: []types.KeySchemaElement{
						{
							AttributeName: aws.String("userId"),
							KeyType:       types.KeyTypeHash,
						},
					},
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
					ProvisionedThroughput: &types.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(5),
						WriteCapacityUnits: aws.Int64(5),
					},
				},
				{
					IndexName: aws.String("OutfitIdIndex"),
					KeySchema: []types.KeySchemaElement{
						{
							AttributeName: aws.String("outfitId"),
							KeyType:       types.KeyTypeHash,
						},
					},
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
					ProvisionedThroughput: &types.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(5),
						WriteCapacityUnits: aws.Int64(5),
					},
				},
			},
		},
		{
			Name: RecommendationsTableName,
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:       types.KeyTypeHash,
				},
			},
			AttributeDef: []types.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("userId"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("outfitId"),
					AttributeType: types.ScalarAttributeTypeS,
				},
			},
			GSIs: []types.GlobalSecondaryIndex{
				{
					IndexName: aws.String("UserIdIndex"),
					KeySchema: []types.KeySchemaElement{
						{
							AttributeName: aws.String("userId"),
							KeyType:       types.KeyTypeHash,
						},
					},
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
					ProvisionedThroughput: &types.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(5),
						WriteCapacityUnits: aws.Int64(5),
					},
				},
				{
					IndexName: aws.String("OutfitIdIndex"),
					KeySchema: []types.KeySchemaElement{
						{
							AttributeName: aws.String("outfitId"),
							KeyType:       types.KeyTypeHash,
						},
					},
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
					ProvisionedThroughput: &types.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(5),
						WriteCapacityUnits: aws.Int64(5),
					},
				},
			},
		},
	}

	for _, table := range tables {
		_, err := client.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
			TableName:              aws.String(table.Name),
			KeySchema:              table.KeySchema,
			AttributeDefinitions:   table.AttributeDef,
			GlobalSecondaryIndexes: table.GSIs,
			ProvisionedThroughput: &types.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(5),
				WriteCapacityUnits: aws.Int64(5),
			},
		})
		if err != nil {
			// If the table already exists, that's fine
			var resourceInUseErr *types.ResourceInUseException
			if ok := errors.As(err, &resourceInUseErr); !ok {
				log.Printf("Error creating table %s: %v", table.Name, err)
				return err
			}
			log.Printf("Table %s already exists", table.Name)
		} else {
			log.Printf("Created table %s", table.Name)
		}
	}

	return nil
}

package dynamodb

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	awsDynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
)

type Table struct {
	Client *awsDynamodb.Client
	Name   string
}

func (t *Table) Create(ctx context.Context, client *awsDynamodb.Client, waitForTable bool) error {
	if t.Client == nil {
		return errors.New("must initialize table with client")
	}
	if t.Name == "" {
		return errors.New("must initialize table with name")
	}
	_, err := client.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   aws.String(t.Name),
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		return errors.Wrap(err, "creating dynamodb table")
	}

	if waitForTable {
		w := dynamodb.NewTableExistsWaiter(client)
		err := w.Wait(ctx,
			&dynamodb.DescribeTableInput{
				TableName: aws.String(t.Name),
			},
			2*time.Minute,
			func(o *dynamodb.TableExistsWaiterOptions) {
				o.MaxDelay = 5 * time.Second
				o.MinDelay = 5 * time.Second
			})
		if err != nil {
			return errors.Wrap(err, "waiting for table to become active")
		}
	}

	return nil
}

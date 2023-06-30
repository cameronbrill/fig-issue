package dynamodb

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsDynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Table struct {
	Client *awsDynamodb.Client
	Name   string
}

func (t *Table) CreateIfNotExists(ctx context.Context) error {
	log.Infof("checking if table exists: %s", t.Name)
	exists, err := t.exists(ctx)
	if err != nil {
		return errors.Wrap(err, "checking if table exists")
	}
	if exists {
		log.Debugf("table %s exists", t.Name)
		return nil
	}
	return errors.Wrap(t.create(ctx, t.Client, true), "creating table")
}

func (t *Table) exists(ctx context.Context) (bool, error) {
	log.Debug("listing tables")
	tables, err := t.Client.ListTables(
		ctx, &awsDynamodb.ListTablesInput{})
	if err != nil {
		return false, errors.Wrap(err, "listing tables")
	}

	log.Debugf("checking if table %s exists among %+v", t.Name, tables.TableNames)
	for _, table := range tables.TableNames {
		if t.Name == table {
			return true, nil
		}
	}
	return false, nil
}

func (t *Table) create(ctx context.Context, client *awsDynamodb.Client, waitForTable bool) error {
	if t.Client == nil {
		return errors.New("must initialize table with client")
	}
	if t.Name == "" {
		return errors.New("must initialize table with name")
	}
	log.Debug("Creating ddb table")
	_, err := client.CreateTable(context.TODO(), &awsDynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("pk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("sk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("pk"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("sk"),
				KeyType:       types.KeyTypeRange,
			},
		},
		TableName:   aws.String(t.Name),
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		return errors.Wrap(err, "creating dynamodb table")
	}

	if waitForTable {
		log.Debug("waiting for ddb table")
		w := awsDynamodb.NewTableExistsWaiter(client)
		const defaultWaitTime = 2 * time.Minute
		const defaultWaitDelay = 5 * time.Second
		err := w.Wait(ctx,
			&awsDynamodb.DescribeTableInput{
				TableName: aws.String(t.Name),
			},
			defaultWaitTime,
			func(o *awsDynamodb.TableExistsWaiterOptions) {
				o.MaxDelay = defaultWaitDelay
				o.MinDelay = defaultWaitDelay
			})
		if err != nil {
			return errors.Wrap(err, "waiting for table to become active")
		}
		log.Debug("finished waiting for ddb table")
	}

	return nil
}

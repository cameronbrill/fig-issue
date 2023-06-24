package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	awsDynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/pkg/errors"
)

func NewClient() (*awsDynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "loading aws config")
	}

	svc := dynamodb.NewFromConfig(cfg)
	return svc, nil
}

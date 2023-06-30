package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	awsDynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/pkg/errors"
)

func NewClient(ctx context.Context) (*awsDynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "loading aws config")
	}

	svc := awsDynamodb.NewFromConfig(cfg)
	return svc, nil
}

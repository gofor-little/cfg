package cfg

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

var (
	// SecretsManagerClient is used to fetch secrets from SecretsManager.
	SecretsManagerClient *secretsmanager.Client
)

// Initialize will initialize the cfg package. Both the profile
// and region parameters are optional if authentication can be achieved
// via another method. For example, environment variables or IAM roles.
func Initialize(ctx context.Context, profile string, region string) error {
	var cfg aws.Config
	var err error

	if profile != "" && region != "" {
		cfg, err = config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile), config.WithRegion(region))
	} else {
		cfg, err = config.LoadDefaultConfig(ctx)
	}
	if err != nil {
		return fmt.Errorf("failed to load default config: %w", err)
	}

	SecretsManagerClient = secretsmanager.NewFromConfig(cfg)

	return nil
}

func checkPackage() error {
	if SecretsManagerClient == nil {
		return errors.New("db.DynamoDBClient is nil, have you called db.Initialize()?")
	}

	return nil
}

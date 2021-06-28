package cfg

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// Load fetches a secret from SecretsManager and parses the value into v using
// json.Unmarshal. To comply with json.Unmarshal's parameter requirements v
// cannot be nil or not a pointer.
func Load(ctx context.Context, secretARN string, v interface{}) error {
	if err := checkPackage(); err != nil {
		return fmt.Errorf("checkPackage call failed: %w", err)
	}

	secret, err := LoadString(ctx, secretARN)
	if err != nil {
		return fmt.Errorf("failed to load secret: %w", err)
	}

	if err := json.Unmarshal([]byte(secret), v); err != nil {
		return fmt.Errorf("failed to unmashal secret value: %w", err)
	}

	return nil
}

// Load fetches a secret from SecretsManager.
func LoadString(ctx context.Context, secretARN string) (string, error) {
	if err := checkPackage(); err != nil {
		return "", fmt.Errorf("checkPackage call failed: %w", err)
	}

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretARN),
	}

	output, err := SecretsManagerClient.GetSecretValue(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to fetch secret: %w", err)
	}

	return *output.SecretString, nil
}

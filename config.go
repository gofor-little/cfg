package cfg

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// Load fetches a secret from SecretsManager and parses the value into v using
// json.Unmarshal. To comply with json.Unmarshal's parameter requirements v
// cannot be nil or not a pointer.
func Load(ctx context.Context, secretArn string, v interface{}) error {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretArn),
	}

	if err := input.Validate(); err != nil {
		return fmt.Errorf("failed to validate input: %w", err)
	}

	output, err := SecretsManagerClient.GetSecretValueWithContext(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to fetch secret: %w", err)
	}

	if err := json.Unmarshal([]byte(*output.SecretString), v); err != nil {
		return fmt.Errorf("failed to unmashal secret value: %w", err)
	}

	return nil
}

package cfg_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/gofor-little/env"
	"github.com/stretchr/testify/require"

	"github.com/gofor-little/cfg"
)

type TestConfig struct {
	Key1 string `json:"key-1"`
	Key2 string `json:"key-2"`
}

func setup(t *testing.T) string {
	if err := env.Load(".env"); err != nil {
		t.Logf("failed to load .env file, ignore if running in CI/CD: %v", err)
	}

	require.NoError(t, cfg.Initialize(context.Background(), env.Get("AWS_PROFILE", ""), env.Get("AWS_REGION", "")))

	data, err := json.Marshal(&TestConfig{"Value-1", "Value-2"})
	require.NoError(t, err)

	input := &secretsmanager.CreateSecretInput{
		Name:         aws.String(fmt.Sprintf("gofor-little-cfg-test-secret-%d", time.Now().Unix())),
		SecretString: aws.String(string(data)),
	}

	output, err := cfg.SecretsManagerClient.CreateSecret(context.Background(), input)
	require.NoError(t, err)

	return *output.ARN
}

func teardown(t *testing.T, secretArn string) {
	input := &secretsmanager.DeleteSecretInput{
		ForceDeleteWithoutRecovery: true,
		SecretId:                   aws.String(secretArn),
	}

	_, err := cfg.SecretsManagerClient.DeleteSecret(context.Background(), input)
	require.NoError(t, err)
}

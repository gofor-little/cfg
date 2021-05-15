package cfg_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/gofor-little/env"
	"github.com/stretchr/testify/require"

	"github.com/gofor-little/cfg"
)

type TestConfig struct {
	Key1 string `json:"key-1"`
	Key2 string `json:"key-2"`
}

func setup(t *testing.T) string {
	if err := env.Load("../../.env"); err != nil {
		t.Logf("failed to load .env file, ignore if running in CI/CD: %v", err)
	}

	require.NoError(t, cfg.Initialize(env.Get("AWS_PROFILE", ""), env.Get("AWS_REGION", "")))

	kmsKeyArn, err := env.MustGet("TEST_KMS_KEY_ARN")
	require.NoError(t, err)

	data, err := json.Marshal(&TestConfig{"Value-1", "Value-2"})
	require.NoError(t, err)

	input := &secretsmanager.CreateSecretInput{
		KmsKeyId:     aws.String(kmsKeyArn),
		Name:         aws.String(fmt.Sprintf("ConfigTest_Integration_%d", time.Now().Unix())),
		SecretString: aws.String(string(data)),
	}
	require.NoError(t, input.Validate())

	output, err := cfg.SecretsManagerClient.CreateSecret(input)
	require.NoError(t, err)

	return *output.ARN
}

func teardown(t *testing.T, secretArn string) {
	input := &secretsmanager.DeleteSecretInput{
		ForceDeleteWithoutRecovery: aws.Bool(true),
		SecretId:                   aws.String(secretArn),
	}
	require.NoError(t, input.Validate())

	_, err := cfg.SecretsManagerClient.DeleteSecret(input)
	require.NoError(t, err)
}

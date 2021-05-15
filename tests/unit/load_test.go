package config_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/gofor-little/aws-sdk-mock"

	"github.com/gofor-little/cfg"
)

func TestLoad(t *testing.T) {
	cfg.SecretsManagerClient = &mock.SecretsManagerClient{}

	testCases := []struct {
		name      string
		secretArn string
		want      error
	}{
		{"TestLoad_Unit", "secret-arn", nil},
	}

	for i, tc := range testCases {
		name := fmt.Sprintf("%s_%d", tc.name, i)

		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.want, cfg.Load(context.Background(), tc.secretArn, &struct {
				Key1 string `json:"key-1"`
				Key2 string `json:"key-2"`
			}{}))
		})
	}
}

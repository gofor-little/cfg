package config_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gofor-little/config"
)

func TestLoad(t *testing.T) {
	secretArn := setup(t)
	defer teardown(t, secretArn)

	testCases := []struct {
		name string
		want error
	}{
		{"TestLoad_Integration", nil},
	}

	for i, tc := range testCases {
		name := fmt.Sprintf("%s_%d", tc.name, i)

		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.want, config.Load(context.Background(), secretArn, &TestConfig{}))
		})
	}
}

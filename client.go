package config

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

var (
	// SecretsManagerClient is used to fetch secrets from SecretsManager.
	SecretsManagerClient secretsmanageriface.SecretsManagerAPI
)

// Initialize will initialize the config package. Both the profile
// and region parameters are optional if authentication can be achieved
// via another method. For example, environment variables or IAM roles.
func Initialize(profile string, region string) error {
	var sess *session.Session
	var err error

	if profile != "" && region != "" {
		sess, err = session.NewSessionWithOptions(session.Options{
			Config: aws.Config{
				Region: aws.String(region),
			},
			Profile: profile,
		})
	} else {
		sess, err = session.NewSession()
	}
	if err != nil {
		return fmt.Errorf("failed to create session.Session: %w", err)
	}

	SecretsManagerClient = secretsmanager.New(sess)

	return nil
}

package vaultsecure

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func TestLoadAWSConfigSupportsSSOSessionProfile(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), "config")
	config := `[profile test-sso]
sso_session = test-session
sso_account_id = 123456789012
sso_role_name = AdministratorAccess
region = us-west-2

[sso-session test-session]
sso_start_url = https://example.awsapps.com/start
sso_region = us-west-2
`
	writeAWSConfig(t, configPath, config)
	configureAWSTestEnvironment(t, configPath, "test-sso")

	cfg, err := loadAWSConfig(context.Background())
	if err != nil {
		t.Fatalf("load modern IAM Identity Center profile: %v", err)
	}
	if cfg.Region != "us-west-2" {
		t.Fatalf("expected profile region us-west-2, got %q", cfg.Region)
	}
	assertCredentialSource(t, cfg.Credentials, aws.CredentialSourceProfileSSO)
}

func TestLoadAWSConfigSupportsAWSLoginProfile(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), "config")
	config := `[profile test-login]
login_session = arn:aws:sts::123456789012:role/AdministratorAccess
region = us-west-2
`
	writeAWSConfig(t, configPath, config)
	configureAWSTestEnvironment(t, configPath, "test-login")
	t.Setenv("AWS_LOGIN_CACHE_DIRECTORY", t.TempDir())

	cfg, err := loadAWSConfig(context.Background())
	if err != nil {
		t.Fatalf("load aws login profile: %v", err)
	}
	if cfg.Region != "us-west-2" {
		t.Fatalf("expected profile region us-west-2, got %q", cfg.Region)
	}
	assertCredentialSource(t, cfg.Credentials, aws.CredentialSourceProfileLogin)
}

func writeAWSConfig(t *testing.T, path, config string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(config), 0o600); err != nil {
		t.Fatalf("write AWS config: %v", err)
	}
}

func configureAWSTestEnvironment(t *testing.T, configPath, profile string) {
	t.Helper()
	t.Setenv("AWS_CONFIG_FILE", configPath)
	t.Setenv("AWS_SHARED_CREDENTIALS_FILE", filepath.Join(t.TempDir(), "credentials"))
	t.Setenv("AWS_PROFILE", profile)
	t.Setenv("AWS_EC2_METADATA_DISABLED", "true")
}

func assertCredentialSource(t *testing.T, provider aws.CredentialsProvider, want aws.CredentialSource) {
	t.Helper()

	sourceProvider, ok := provider.(aws.CredentialProviderSource)
	if !ok {
		t.Fatalf("credentials provider %T does not report its source", provider)
	}
	for _, source := range sourceProvider.ProviderSources() {
		if source == want {
			return
		}
	}

	t.Fatalf("expected credential source %v, got %v", want, sourceProvider.ProviderSources())
}

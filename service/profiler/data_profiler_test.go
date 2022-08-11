package profiler

import (
	"testing"

	awsmiscclient "github.com/STollenaar/aws-misc-client"
)

// TestListProfiles defined data resource for the terraform plugin
func TestListProfiles(t *testing.T) {
	client, err := awsmiscclient.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	profiles, err := client.Profiler.GetProfiles()
	if err != nil {
		t.Fatal(err)
	}

	if len(profiles) == 0 {
		t.Fatal("Error, not profiles found or none configured\n")
	}
	t.Logf("Received %d profile(s)\n", len(profiles))

	for _, profile := range profiles {
		if profile.Name == "" {
			t.Fatal("Error, Profile name is not defined\n")
		}
		if profile.AWSAccessKeyId == "" {
			t.Fatal("Error, AccesKeyId is not defined\n")
		}
		if profile.AWSSecretAccessKey == "" {
			t.Fatal("Error, SecretAccesKeyId is not defined\n")
		}
	}
}

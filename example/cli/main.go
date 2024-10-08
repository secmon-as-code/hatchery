package main

import (
	"os"

	"github.com/secmon-as-code/hatchery"
	"github.com/secmon-as-code/hatchery/destination/gcs"
	"github.com/secmon-as-code/hatchery/pkg/types/secret"
	"github.com/secmon-as-code/hatchery/source/slack"
)

func main() {
	streams := []*hatchery.Stream{
		hatchery.NewStream(
			// Source: Slack Audit API
			slack.New(secret.NewString(os.Getenv("SLACK_TOKEN"))),
			// Destination: Google Cloud Storage, bucket name is "mizutani-test"
			gcs.New("mizutani-test"),

			// With ID
			hatchery.WithID("slack-to-gcs"),
		),
	}

	// You can run CLI with args such as `go run main.go -s slack-to-gcs`
	if err := hatchery.New(streams).CLI(os.Args); err != nil {
		panic(err)
	}
}

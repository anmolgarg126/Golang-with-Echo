package config

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func configureSentry() {
	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: "https://a2bdf73fad2a4c758b6c088491a7301c@o996077.ingest.sentry.io/4504133443584000",
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		// Environment: "DEV",
		// Release:     "my-project-name@1.0.0",
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		// Debug:            true,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")
}

package config

// Config is the configuration for the application.
// Its values are loaded from environment variables using
// the envconfig package in the init of the main package.
type Config struct {
	ProjectID string `envconfig:"GCP_PROJECT_ID" required:"true"`
	Region    string `envconfig:"GCP_REGION" required:"true"`
	Port      string `envconfig:"PORT" default:"8080"`
	Debug     bool   `envconfig:"DEBUG" default:"false"`
}

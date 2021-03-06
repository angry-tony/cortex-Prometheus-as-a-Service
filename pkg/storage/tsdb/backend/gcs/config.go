package gcs

import (
	"flag"
)

// Config holds the config options for GCS backend
type Config struct {
	BucketName     string `yaml:"bucket_name"`
	ServiceAccount string `yaml:"service_account"`
}

// RegisterFlags registers the flags for TSDB GCS storage
func (cfg *Config) RegisterFlags(f *flag.FlagSet) {
	f.StringVar(&cfg.BucketName, "experimental.tsdb.gcs.bucket-name", "", "GCS bucket name")
	f.StringVar(&cfg.ServiceAccount, "experimental.tsdb.gcs.service-account", "", "JSON representing either a Google Developers Console client_credentials.json file or a Google Developers service account key file. If empty, fallback to Google default logic.")
}
